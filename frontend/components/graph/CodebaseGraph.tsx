'use client';

import { useCallback, useEffect } from 'react';
import ReactFlow, {
  Node,
  Edge,
  Controls,
  Background,
  MiniMap,
  useNodesState,
  useEdgesState,
  addEdge,
  Connection,
  BackgroundVariant,
} from 'reactflow';
import 'reactflow/dist/style.css';
import type { Node as AnalysisNode, Edge as AnalysisEdge } from '@/types/analysis';

interface CodebaseGraphProps {
  nodes: AnalysisNode[];
  edges: AnalysisEdge[];
  selectedNodeId?: string | null;
  onNodeClick?: (node: AnalysisNode) => void;
}

// Convert our analysis nodes to ReactFlow nodes
const convertToFlowNodes = (analysisNodes: AnalysisNode[], dependentNodeIds: Set<string> = new Set()): Node[] => {
  return analysisNodes.map((node, index) => {
    // Extract filename from full path (handles both / and \ separators)
    const filename = node.id.split(/[\/\\]/).pop() || node.id;
    const isDependent = dependentNodeIds.has(node.id);
    
    return {
      id: node.id,
      type: 'default',
      position: {
        x: (index % 5) * 250,
        y: Math.floor(index / 5) * 150,
      },
      data: {
        label: filename,
        fullPath: node.id,
        ...node.metrics,
      },
      style: {
        background: getRiskColor(node.metrics.risk_level),
        color: '#fff',
        border: isDependent ? '3px solid #fbbf24' : '2px solid #222',
        borderRadius: '8px',
        padding: '10px',
        fontSize: '12px',
        fontWeight: 500,
        boxShadow: isDependent ? '0 0 10px rgba(251, 191, 36, 0.5)' : 'none',
      },
    };
  });
};

// Convert our analysis edges to ReactFlow edges
const convertToFlowEdges = (analysisEdges: AnalysisEdge[]): Edge[] => {
  return analysisEdges.map((edge, index) => ({
    id: `edge-${index}`,
    source: edge.source,
    target: edge.target,
    type: 'smoothstep',
    animated: true,
    label: edge.type,
    style: { stroke: '#888' },
  }));
};

// Get color based on risk level
function getRiskColor(riskLevel: string): string {
  switch (riskLevel) {
    case 'high':
      return '#ef4444'; // red
    case 'medium':
      return '#f59e0b'; // orange
    case 'low':
      return '#10b981'; // green
    default:
      return '#6b7280'; // gray
  }
}

export default function CodebaseGraph({ nodes, edges, selectedNodeId, onNodeClick }: CodebaseGraphProps) {
  const [flowNodes, setFlowNodes, onNodesChange] = useNodesState(convertToFlowNodes(nodes));
  const [flowEdges, setFlowEdges, onEdgesChange] = useEdgesState(convertToFlowEdges(edges));

  // Calculate dependent nodes when selectedNodeId changes
  useEffect(() => {
    const dependentNodeIds = new Set<string>();
    if (selectedNodeId) {
      // Find all nodes that import the selected node (reverse edges)
      edges.forEach((edge) => {
        if (edge.source === selectedNodeId) {
          dependentNodeIds.add(edge.target);
        }
      });
    }
    setFlowNodes(convertToFlowNodes(nodes, dependentNodeIds));
  }, [selectedNodeId, nodes, edges, setFlowNodes]);

  // Sync analysis data to React Flow state when it changes
  useEffect(() => {
    setFlowEdges(convertToFlowEdges(edges));
  }, [edges, setFlowEdges]);

  const onConnect = useCallback(
    (params: Connection) => setFlowEdges((eds) => addEdge(params, eds)),
    [setFlowEdges]
  );

  const handleNodeClick = useCallback(
    (_event: React.MouseEvent, node: Node) => {
      const originalNode = nodes.find((n) => n.id === node.id);
      if (originalNode && onNodeClick) {
        onNodeClick(originalNode);
      }
    },
    [nodes, onNodeClick]
  );

  if (nodes.length === 0) {
    return (
      <div className="flex items-center justify-center h-full bg-gray-50 dark:bg-gray-900">
        <div className="text-center">
          <p className="text-gray-500 dark:text-gray-400 text-lg">
            No data to visualize yet
          </p>
          <p className="text-gray-400 dark:text-gray-500 text-sm mt-2">
            Enter a repository URL to analyze
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="w-full h-full">
      <ReactFlow
        nodes={flowNodes}
        edges={flowEdges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        onConnect={onConnect}
        onNodeClick={handleNodeClick}
        fitView
        attributionPosition="bottom-left"
      >
        <Controls />
        <MiniMap
          nodeColor={(node) => {
            const riskLevel = node.data?.risk_level || 'low';
            return getRiskColor(riskLevel);
          }}
          maskColor="rgba(0, 0, 0, 0.1)"
        />
        <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
      </ReactFlow>
    </div>
  );
}