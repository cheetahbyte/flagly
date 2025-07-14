import "./index.css";

import { HeadRow } from "./components/HeadRow";
import AppHeader from "./components/Header";
import FeatureList from "./components/FeatureList";
import { SidebarProvider } from "./components/ui/sidebar";
import { AppSidebar } from "./components/AppSidebar";
import type { Flag, Environment } from "@/lib/types";

import { useEffect, useState } from "react";
import { useApiClient } from "./hooks/useApiClient";
export function App() {
  const [flags, setFlags] = useState<Flag[]>([]);
  const [environments, setEnvironments] = useState<string[]>([]);
  const apiClient = useApiClient();

  useEffect(() => {
    const fetchFlags = async () => {
      const res = await apiClient.get<Flag[]>("/flags");
      setFlags(res);
    };
    const fetchEnvironments = async () => {
      const res = await apiClient.get<string[]>("/environments");
      setEnvironments(res);
    };
    fetchFlags();
    fetchEnvironments();
  }, []);
  return (
    <SidebarProvider>
      <div className="container mx-auto p-8 text-center relative z-10">
        <AppSidebar />
        <AppHeader />
        <HeadRow flags={flags} environments={environments} />
        <FeatureList flags={flags} environments={environments} />
      </div>
    </SidebarProvider>
  );
}

export default App;
