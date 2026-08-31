'use client';

import { useState } from 'react';
import { Search, Loader2, AlertCircle, Code2 } from 'lucide-react';
import CodebaseGraph from './graph/CodebaseGraph';
import Sidebar from './ui/Sidebar';
import { api } from '@/lib/api/client';
import type { AnalysisResult, Node } from '@/types/analysis';

export default function Dashboard() {
  const [repoUrl, setRepoUrl] = useState('');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [analysisResult, setAnalysisResult] = useState<AnalysisResult | null>(null);
  const [selectedNode, setSelectedNode] = useState<Node | null>(null);

  const handleAnalyze = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!repoUrl.trim()) {
      setError('Please enter a repository URL');
      return;
    }

    setLoading(true);
    setError(null);
    setSelectedNode(null);

    try {
      const result = await api.analyzeRepository(repoUrl);
      setAnalysisResult(result);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to analyze repository');
      console.error('Analysis error:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleNodeClick = (node: Node) => {
    setSelectedNode(node);
  };

  const handleCloseSidebar = () => {
    setSelectedNode(null);
  };

  return (
    <div className="flex flex-col h-screen bg-gray-50 dark:bg-gray-950">
      {/* Header */}
      <header className="bg-white dark:bg-gray-900 border-b border-gray-200 dark:border-gray-800 px-6 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="w-10 h-10 bg-linear-to-br from-blue-500 to-purple-600 rounded-lg flex items-center justify-center">
              <span className="text-white font-bold text-xl">D</span>
            </div>
            <div>
              <h1 className="text-2xl font-bold text-gray-900 dark:text-white">
                DevLens
              </h1>
              <p className="text-sm text-gray-500 dark:text-gray-400">
                Codebase Intelligence Platform
              </p>
            </div>
          </div>

          {/* Summary Stats */}
          {analysisResult && (
            <div className="flex gap-6">
              <div className="text-center">
                <div className="text-2xl font-bold text-gray-900 dark:text-white">
                  {analysisResult.summary.total_files}
                </div>
                <div className="text-xs text-gray-500 dark:text-gray-400">Files</div>
              </div>
              <div className="text-center">
                <div className="text-2xl font-bold text-gray-900 dark:text-white">
                  {analysisResult.summary.total_functions}
                </div>
                <div className="text-xs text-gray-500 dark:text-gray-400">Functions</div>
              </div>
              <div className="text-center">
                <div className="text-2xl font-bold text-red-600 dark:text-red-400">
                  {analysisResult.summary.high_risk_files}
                </div>
                <div className="text-xs text-gray-500 dark:text-gray-400">High Risk</div>
              </div>
            </div>
          )}
        </div>

        {/* Search Bar */}
        <form onSubmit={handleAnalyze} className="mt-4">
          <div className="flex gap-2">
            <div className="flex-1 relative">
              <Code2 className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
              <input
                type="text"
                value={repoUrl}
                onChange={(e) => setRepoUrl(e.target.value)}
                placeholder="Enter GitHub repository URL (e.g., https://github.com/user/repo)"
                className="w-full pl-10 pr-4 py-3 border border-gray-300 dark:border-gray-700 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent bg-white dark:bg-gray-800 text-gray-900 dark:text-white placeholder-gray-400 dark:placeholder-gray-500"
                disabled={loading}
              />
            </div>
            <button
              type="submit"
              disabled={loading}
              className="px-6 py-3 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              {loading ? (
                <>
                  <Loader2 className="w-5 h-5 animate-spin" />
                  Analyzing...
                </>
              ) : (
                <>
                  <Search className="w-5 h-5" />
                  Analyze
                </>
              )}
            </button>
          </div>
        </form>

        {/* Error Message */}
        {error && (
          <div className="mt-3 p-3 bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-lg flex items-center gap-2">
            <AlertCircle className="w-5 h-5 text-red-600 dark:text-red-400" />
            <p className="text-sm text-red-800 dark:text-red-200">{error}</p>
          </div>
        )}

        {/* Info Message */}
        {analysisResult?.message && (
          <div className="mt-3 p-3 bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800 rounded-lg">
            <p className="text-sm text-blue-800 dark:text-blue-200">{analysisResult.message}</p>
          </div>
        )}
      </header>

      {/* Main Content */}
      <div className="flex-1 flex overflow-hidden">
        {/* Graph Visualization */}
        <div className="flex-1">
          <CodebaseGraph
            nodes={analysisResult?.nodes || []}
            edges={analysisResult?.edges || []}
            selectedNodeId={selectedNode?.id || null}
            onNodeClick={handleNodeClick}
          />
        </div>

        {/* Sidebar */}
        {selectedNode && (
          <Sidebar selectedNode={selectedNode} edges={analysisResult?.edges || []} onClose={handleCloseSidebar} />
        )}
      </div>
    </div>
  );
}