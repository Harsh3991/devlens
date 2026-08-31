package parser

import (
 "fmt"
 "os"
 "path/filepath"
 "strings"

 tree_sitter "github.com/tree-sitter/go-tree-sitter"
 tree_sitter_javascript "github.com/tree-sitter/tree-sitter-javascript/bindings/go"

 "devlens-backend/internal/models"
)

// ASTParser handles parsing of source files using Tree-sitter
type ASTParser struct {
 parser *tree_sitter.Parser
}

// NewASTParser creates a new AST parser instance
func NewASTParser() *ASTParser {
 parser := tree_sitter.NewParser()
 return &ASTParser{
 parser: parser,
 }
}

// Close cleans up parser resources
func (p *ASTParser) Close() {
 if p.parser != nil {
 p.parser.Close()
 }
}

// ParseFile parses a single file and extracts metadata
func (p *ASTParser) ParseFile(filePath string) (*models.FileInfo, error) {
// Read file content
 content, err := os.ReadFile(filePath)
 if err != nil {
 return nil, fmt.Errorf("failed to read file: %w", err)
 }

// Set language based on file extension
 ext := filepath.Ext(filePath)
 if err := p.setLanguage(ext); err != nil {
 return nil, fmt.Errorf("unsupported file type %s: %w", ext, err)
 }

// Parse the file
 tree := p.parser.Parse(content, nil)
 if tree == nil {
 return nil, fmt.Errorf("failed to parse file")
 }
 defer tree.Close()

 root := tree.RootNode()

// Extract file information
 fileInfo := &models.FileInfo{
 Path: filePath,
 Imports: []string{},
 Exports: []string{},
 Functions: []models.FunctionInfo{},
 Complexity: 0,
 LinesOfCode: int(root.EndPosition().Row) + 1,
 }

// Traverse AST and extract information
 p.traverseNode(root, content, fileInfo)

 return fileInfo, nil
}

// setLanguage sets the appropriate Tree-sitter language for the parser
func (p *ASTParser) setLanguage(ext string) error {
 var language *tree_sitter.Language

 switch ext {
 case ".js", ".mjs", ".cjs", ".jsx":
 language = tree_sitter.NewLanguage(tree_sitter_javascript.Language())
 case ".ts", ".tsx":
// Use JavaScript parser for TypeScript files
// The JavaScript parser can handle most TypeScript syntax
 language = tree_sitter.NewLanguage(tree_sitter_javascript.Language())
 default:
 return fmt.Errorf("unsupported file extension: %s", ext)
 }

 p.parser.SetLanguage(language)
 return nil
}

// traverseNode recursively traverses the AST and extracts information
func (p *ASTParser) traverseNode(node *tree_sitter.Node, content []byte, fileInfo *models.FileInfo) {
 if node == nil {
 return
 }

 nodeType := node.Kind()

// Extract imports
 if nodeType == "import_statement" || nodeType == "import_declaration" {
 importPath := p.extractImportPath(node, content)
 if importPath != "" {
 fileInfo.Imports = append(fileInfo.Imports, importPath)
 }
 }

// Extract exports
 if strings.HasPrefix(nodeType, "export_") {
// Mark that this file has exports
 if !contains(fileInfo.Exports, fileInfo.Path) {
 fileInfo.Exports = append(fileInfo.Exports, fileInfo.Path)
 }
 }

// Extract function declarations
 if nodeType == "function_declaration" || nodeType == "function" ||
 nodeType == "arrow_function" || nodeType == "method_definition" {
 funcInfo := p.extractFunctionInfo(node, content)
 if funcInfo != nil {
 fileInfo.Functions = append(fileInfo.Functions, *funcInfo)
 fileInfo.Complexity += funcInfo.Complexity
 }
 }

// Calculate complexity for control flow statements
 if p.isComplexityNode(nodeType) {
 fileInfo.Complexity++
 }

// Recursively traverse children
 for i := uint(0); i < node.ChildCount(); i++ {
 child := node.Child(i)
 if child != nil {
 p.traverseNode(child, content, fileInfo)
 }
 }
}

// extractImportPath extracts the import path from an import statement
func (p *ASTParser) extractImportPath(node *tree_sitter.Node, content []byte) string {
// Look for string literal in import statement
 for i := uint(0); i < node.ChildCount(); i++ {
 child := node.Child(i)
 if child != nil && child.Kind() == "string" {
// Extract the string content
 text := child.Utf8Text(content)
// Remove quotes
 text = strings.Trim(text, `"'`)
 return text
 }
 }
 return ""
}

// extractFunctionInfo extracts information about a function
func (p *ASTParser) extractFunctionInfo(node *tree_sitter.Node, content []byte) *models.FunctionInfo {
 funcInfo := &models.FunctionInfo{
 LineStart: int(node.StartPosition().Row) + 1,
 LineEnd: int(node.EndPosition().Row) + 1,
 Complexity: 1, // Base complexity
 }

// Try to extract function name
 for i := uint(0); i < node.ChildCount(); i++ {
 child := node.Child(i)
 if child != nil && child.Kind() == "identifier" {
 funcInfo.Name = child.Utf8Text(content)
 break
 }
 }

// If no name found, use anonymous
 if funcInfo.Name == "" {
 funcInfo.Name = "anonymous"
 }

// Calculate cyclomatic complexity for this function
 funcInfo.Complexity += p.calculateFunctionComplexity(node)

 return funcInfo
}

// calculateFunctionComplexity calculates cyclomatic complexity for a function
func (p *ASTParser) calculateFunctionComplexity(node *tree_sitter.Node) int {
 complexity := 0

 if node == nil {
 return complexity
 }

 nodeType := node.Kind()
 if p.isComplexityNode(nodeType) {
 complexity++
 }

// Recursively check children
 for i := uint(0); i < node.ChildCount(); i++ {
 child := node.Child(i)
 if child != nil {
 complexity += p.calculateFunctionComplexity(child)
 }
 }

 return complexity
}

// isComplexityNode checks if a node type contributes to cyclomatic complexity
func (p *ASTParser) isComplexityNode(nodeType string) bool {
 complexityNodes := map[string]bool{
 "if_statement": true,
 "else_clause": true,
 "switch_statement": true,
 "case_clause": true,
 "for_statement": true,
 "for_in_statement": true,
 "while_statement": true,
 "do_statement": true,
 "catch_clause": true,
 "ternary_expression": true,
 "binary_expression": false, // We'll check for && and ||
 }

 return complexityNodes[nodeType]
}

// Helper function to check if a string slice contains a value
func contains(slice []string, item string) bool {
 for _, s := range slice {
 if s == item {
 return true
 }
 }
 return false
}