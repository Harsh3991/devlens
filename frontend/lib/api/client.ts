import axios from 'axios';
import type { AnalysisResult, AnalyzeRequest, HealthCheckResponse } from '@/types/analysis';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

// Create axios instance with default config
const apiClient = axios.create({
 baseURL: API_URL,
 headers: {
 'Content-Type': 'application/json',
 },
 timeout: 60000, // 60 seconds for analysis requests
});

// API functions
export const api = {
// Health check
 healthCheck: async (): Promise<HealthCheckResponse> => {
const response = await apiClient.get<HealthCheckResponse>('/health');
return response.data;
 },

// Analyze repository
 analyzeRepository: async (sourceUrl: string): Promise<AnalysisResult> => {
const payload: AnalyzeRequest = { source_url: sourceUrl };
const response = await apiClient.post<AnalysisResult>('/api/v1/analyze', payload);
return response.data;
 },
};

export default apiClient;