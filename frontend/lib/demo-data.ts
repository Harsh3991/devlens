// Demo data for testing the visualization without backend
import type { AnalysisResult } from '@/types/analysis';

export const demoAnalysisResult: AnalysisResult = {
 repository: 'demo/sample-project',
 summary: {
 total_files: 8,
 total_functions: 24,
 high_risk_files: 2,
 },
 nodes: [
 {
 id: 'src/index.ts',
 type: 'file',
 metrics: {
 complexity: 5,
 functions_count: 3,
 risk_level: 'low',
 lines_of_code: 120,
 },
 },
 {
 id: 'src/auth.ts',
 type: 'file',
 metrics: {
 complexity: 18,
 functions_count: 6,
 risk_level: 'high',
 lines_of_code: 280,
 },
 },
 {
 id: 'src/api/users.ts',
 type: 'file',
 metrics: {
 complexity: 12,
 functions_count: 4,
 risk_level: 'medium',
 lines_of_code: 195,
 },
 },
 {
 id: 'src/api/posts.ts',
 type: 'file',
 metrics: {
 complexity: 8,
 functions_count: 3,
 risk_level: 'low',
 lines_of_code: 145,
 },
 },
 {
 id: 'src/utils/validation.ts',
 type: 'file',
 metrics: {
 complexity: 15,
 functions_count: 5,
 risk_level: 'high',
 lines_of_code: 220,
 },
 },
 {
 id: 'src/utils/helpers.ts',
 type: 'file',
 metrics: {
 complexity: 6,
 functions_count: 8,
 risk_level: 'low',
 lines_of_code: 160,
 },
 },
 {
 id: 'src/config.ts',
 type: 'file',
 metrics: {
 complexity: 2,
 functions_count: 1,
 risk_level: 'low',
 lines_of_code: 45,
 },
 },
 {
 id: 'src/types.ts',
 type: 'file',
 metrics: {
 complexity: 1,
 functions_count: 0,
 risk_level: 'low',
 lines_of_code: 85,
 },
 },
 ],
 edges: [
 {
 source: 'src/index.ts',
 target: 'src/auth.ts',
 type: 'import',
 },
 {
 source: 'src/index.ts',
 target: 'src/config.ts',
 type: 'import',
 },
 {
 source: 'src/auth.ts',
 target: 'src/utils/validation.ts',
 type: 'import',
 },
 {
 source: 'src/auth.ts',
 target: 'src/types.ts',
 type: 'import',
 },
 {
 source: 'src/api/users.ts',
 target: 'src/auth.ts',
 type: 'import',
 },
 {
 source: 'src/api/users.ts',
 target: 'src/utils/helpers.ts',
 type: 'import',
 },
 {
 source: 'src/api/posts.ts',
 target: 'src/auth.ts',
 type: 'import',
 },
 {
 source: 'src/api/posts.ts',
 target: 'src/utils/helpers.ts',
 type: 'import',
 },
 {
 source: 'src/utils/validation.ts',
 target: 'src/types.ts',
 type: 'import',
 },
 ],
};