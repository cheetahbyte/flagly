import { ApiClient } from "@/lib/api";
import { useMemo } from "react";

const apiClient = new ApiClient();

export function useApiClient() {
  return useMemo(() => apiClient, []);
}
