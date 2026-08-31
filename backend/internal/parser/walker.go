package parser

import (
 "io/fs"
 "path/filepath"
 "strings"
 "sync"
)

// SupportedExtensions defines which file types we can parse
var SupportedExtensions = map[string]bool{
 ".ts": true,
 ".tsx": true,
 ".js": true,
 ".jsx": true,
 ".go": true,
 ".py": true,
}

// FileWalker handles concurrent directory traversal
type FileWalker struct {
 rootPath string
 files []string
 mu sync.Mutex
}

// NewFileWalker creates a new FileWalker instance
func NewFileWalker(rootPath string) *FileWalker {
 return &FileWalker{
 rootPath: rootPath,
 files: make([]string, 0),
 }
}

// Walk traverses the directory and collects all supported files
func (fw *FileWalker) Walk() ([]string, error) {
 err := filepath.WalkDir(fw.rootPath, func(path string, d fs.DirEntry, err error) error {
 if err != nil {
 return err
 }

// Skip hidden directories and node_modules
 if d.IsDir() {
 name := d.Name()
 if strings.HasPrefix(name, ".") || name == "node_modules" || name == "vendor" || name == "dist" || name == "build" {
 return filepath.SkipDir
 }
 return nil
 }

// Check if file has supported extension
 ext := filepath.Ext(path)
 if SupportedExtensions[ext] {
 fw.mu.Lock()
 fw.files = append(fw.files, path)
 fw.mu.Unlock()
 }

 return nil
 })

 return fw.files, err
}

// GetFiles returns the collected files
func (fw *FileWalker) GetFiles() []string {
 fw.mu.Lock()
 defer fw.mu.Unlock()
 return fw.files
}