package parser

import (
 "fmt"
 "os"
 "os/exec"
 "path/filepath"
 "strings"
)

// RepoCloner handles cloning of Git repositories
type RepoCloner struct {
 tempDir string
}

// NewRepoCloner creates a new repository cloner
func NewRepoCloner() *RepoCloner {
 return &RepoCloner{
 tempDir: os.TempDir(),
 }
}

// CloneRepository clones a Git repository to a temporary directory
func (rc *RepoCloner) CloneRepository(repoURL string) (string, error) {
// Validate URL
 if !strings.HasPrefix(repoURL, "http://") && !strings.HasPrefix(repoURL, "https://") {
 return "", fmt.Errorf("invalid repository URL: must start with http:// or https://")
 }

// Create a unique temporary directory
 repoName := rc.extractRepoName(repoURL)
 cloneDir := filepath.Join(rc.tempDir, "devlens", repoName)

// Remove existing directory if it exists
 if err := os.RemoveAll(cloneDir); err != nil {
 return "", fmt.Errorf("failed to clean up existing directory: %w", err)
 }

// Create parent directory
 if err := os.MkdirAll(filepath.Dir(cloneDir), 0755); err != nil {
 return "", fmt.Errorf("failed to create temp directory: %w", err)
 }

// Clone the repository
 cmd := exec.Command("git", "clone", "--depth", "1", repoURL, cloneDir)
 output, err := cmd.CombinedOutput()
 if err != nil {
 return "", fmt.Errorf("failed to clone repository: %w\nOutput: %s", err, string(output))
 }

 return cloneDir, nil
}

// CleanupRepository removes the cloned repository directory
func (rc *RepoCloner) CleanupRepository(repoPath string) error {
 if repoPath == "" || repoPath == "/" {
 return fmt.Errorf("invalid repository path")
 }

// Only remove if it's in our temp directory
 if !strings.HasPrefix(repoPath, filepath.Join(rc.tempDir, "devlens")) {
 return fmt.Errorf("refusing to delete directory outside temp area")
 }

 return os.RemoveAll(repoPath)
}

// extractRepoName extracts the repository name from a URL
func (rc *RepoCloner) extractRepoName(repoURL string) string {
// Remove .git suffix if present
 repoURL = strings.TrimSuffix(repoURL, ".git")

// Extract the last part of the URL
 parts := strings.Split(repoURL, "/")
 if len(parts) > 0 {
 return parts[len(parts)-1]
 }

 return "unknown-repo"
}

// IsGitInstalled checks if git is available in the system
func IsGitInstalled() bool {
 cmd := exec.Command("git", "--version")
 return cmd.Run() == nil
}