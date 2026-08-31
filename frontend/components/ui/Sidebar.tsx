'use client';

import { X, FileCode, AlertCircle, TrendingUp, Activity } from 'lucide-react';
import type { Node, Edge } from '@/types/analysis';

interface SidebarProps {
  selectedNode: Node | null;
  edges?: Edge[];
  onClose: () => void;
}

export default function Sidebar({ selectedNode, edges = [], onClose }: SidebarProps) {
  if (!selectedNode) return null;

  const { id, metrics } = selectedNode;
  const fileName = id.split(/[\/\\]/).pop() || id;

  // Calculate impact analysis - files that depend on this file
  const dependentFiles = edges
    .filter((edge) => edge.source === id)
    .map((edge) => edge.target)
    .map((targetId) => targetId.split(/[\/\\]/).pop() || targetId);

  const getRiskBadgeColor = (risk: string) => {
    switch (risk) {
      case 'high':
        return 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200';
      case 'medium':
        return 'bg-orange-100 text-orange-800 dark:bg-orange-900 dark:text-orange-200';
      case 'low':
        return 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200';
      default:
        return 'bg-gray-100 text-gray-800 dark:bg-gray-900 dark:text-gray-200';
    }
  };

  return (
    <div className="w-96 h-full bg-white dark:bg-gray-800 border-l border-gray-200 dark:border-gray-700 flex flex-col">
      {/* Header */}
      <div className="flex items-center justify-between p-4 border-b border-gray-200 dark:border-gray-700">
        <h2 className="text-lg font-semibold text-gray-900 dark:text-white flex items-center gap-2">
          <FileCode className="w-5 h-5" />
          File Details
        </h2>
        <button
          onClick={onClose}
          className="p-1 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md transition-colors"
          aria-label="Close sidebar"
        >
          <X className="w-5 h-5 text-gray-500 dark:text-gray-400" />
        </button>
      </div>

      {/* Content */}
      <div className="flex-1 overflow-y-auto p-4 space-y-6">
        {/* File Name */}
        <div>
          <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400 mb-1">
            File Name
          </h3>
          <p className="text-base font-mono text-gray-900 dark:text-white break-all">
            {fileName}
          </p>
          <p className="text-xs text-gray-500 dark:text-gray-400 mt-1 font-mono break-all">
            {id}
          </p>
        </div>

        {/* Risk Level */}
        <div>
          <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">
            Risk Level
          </h3>
          <span
            className={`inline-flex items-center gap-1 px-3 py-1 rounded-full text-sm font-medium ${getRiskBadgeColor(
              metrics.risk_level
            )}`}
          >
            <AlertCircle className="w-4 h-4" />
            {metrics.risk_level.toUpperCase()}
          </span>
        </div>

        {/* Metrics */}
        <div className="space-y-4">
          <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400">
            Metrics
          </h3>

          {/* Complexity */}
          <div className="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-900 rounded-lg">
            <div className="flex items-center gap-2">
              <Activity className="w-5 h-5 text-blue-500" />
              <span className="text-sm font-medium text-gray-700 dark:text-gray-300">
                Cyclomatic Complexity
              </span>
            </div>
            <span className="text-lg font-bold text-gray-900 dark:text-white">
              {metrics.complexity}
            </span>
          </div>

          {/* Functions Count */}
          <div className="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-900 rounded-lg">
            <div className="flex items-center gap-2">
              <TrendingUp className="w-5 h-5 text-purple-500" />
              <span className="text-sm font-medium text-gray-700 dark:text-gray-300">
                Functions
              </span>
            </div>
            <span className="text-lg font-bold text-gray-900 dark:text-white">
              {metrics.functions_count}
            </span>
          </div>

          {/* Lines of Code */}
          {metrics.lines_of_code !== undefined && (
            <div className="flex items-center justify-between p-3 bg-gray-50 dark:bg-gray-900 rounded-lg">
              <div className="flex items-center gap-2">
                <FileCode className="w-5 h-5 text-green-500" />
                <span className="text-sm font-medium text-gray-700 dark:text-gray-300">
                  Lines of Code
                </span>
              </div>
              <span className="text-lg font-bold text-gray-900 dark:text-white">
                {metrics.lines_of_code}
              </span>
            </div>
          )}
        </div>

        {/* Impact Analysis Section */}
        <div>
          <h3 className="text-sm font-medium text-gray-500 dark:text-gray-400 mb-2">
            Impact Analysis
          </h3>
          {dependentFiles.length > 0 ? (
            <div className="p-4 bg-amber-50 dark:bg-amber-900/20 border border-amber-200 dark:border-amber-800 rounded-lg">
              <p className="text-sm font-medium text-amber-800 dark:text-amber-200 mb-3">
                {dependentFiles.length} file{dependentFiles.length !== 1 ? 's' : ''} depend on this:
              </p>
              <div className="space-y-2 max-h-48 overflow-y-auto">
                {dependentFiles.map((file, idx) => (
                  <div
                    key={idx}
                    className="p-2 bg-white dark:bg-gray-700 rounded border border-amber-200 dark:border-amber-700 text-xs text-gray-700 dark:text-gray-300 font-mono break-all hover:bg-amber-50 dark:hover:bg-gray-600 transition-colors"
                  >
                    {file}
                  </div>
                ))}
              </div>
              <p className="text-xs text-amber-600 dark:text-amber-300 mt-3">
                ⚠️ Changes to this file may affect {dependentFiles.length} other file{dependentFiles.length !== 1 ? 's' : ''}.
              </p>
            </div>
          ) : (
            <div className="p-4 bg-gray-50 dark:bg-gray-900 border border-gray-200 dark:border-gray-700 rounded-lg">
              <p className="text-sm text-gray-600 dark:text-gray-400">
                No files depend on this file (no blast radius).
              </p>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}