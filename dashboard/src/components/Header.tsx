import React, { useEffect, useState } from "react";
import logo from "@/logo-dark.svg";
import { SidebarTrigger } from "./ui/sidebar";
import { useApiClient } from "@/hooks/useApiClient";
import type { ApiStatusResponse } from "@/lib/types";

export default function AppHeader() {
  const apiClient = useApiClient();
  const [version, setVersion] = useState<string>("");

  useEffect(() => {
    const fetchServerState = async () => {
      const res = await apiClient.get<ApiStatusResponse>("/status");
      setVersion(res.version);
    };
    fetchServerState();
  }, []);

  return (
    <header className="flex items-center justify-between p-4 border-b border-gray-200 mb-6">
      <div className="flex items-center">
        <SidebarTrigger />
      </div>

      <nav className="flex space-x-4">
        <a
          href="https://github.com/cheetahbyte/flagly"
          className="text-sm font-medium text-gray-700 hover:text-gray-900"
        >
          Github
        </a>
        <p className="text-sm font-medium text-gray-700 hover:text-gray-900">
          {version}
        </p>
      </nav>
    </header>
  );
}
