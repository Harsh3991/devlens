// Type definitions matching the backend API response

export interface AnalysisResult {
 repository: string;
 summary: Summary;
 nodes: Node[];
 edges: Edge[];
 message?: string;
}

export interface Summary {
 total_files: number;
 total_functions: number;
 high_risk_files: number;
}

export interface Node {
 id: string;
 type: string;
 metrics: Metrics;
}

export interface Metrics {
 complexity: number;
 functions_count: number;
 risk_level: 'low' | 'medium' | 'high';
 lines_of_code?: number;
}

export interface Edge {
 source: string;
 target: string;
 type: string;
}

export interface AnalyzeRequest {
 source_url: string;
}

export interface HealthCheckResponse {
 status: string;
 service: string;
 timestamp: string;
 version: string;
}