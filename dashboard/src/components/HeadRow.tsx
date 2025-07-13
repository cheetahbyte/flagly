import StatCard from "./StatCard";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import {
  Eye,
  Copy,
  ExternalLink,
  TrendingUp,
  TrendingDown,
  Minus,
} from "lucide-react";
export function HeadRow() {
  const totalFlags = 10;
  return (
    <div className="grid gap-4 md:grid-cols-4">
      <StatCard
        title="Total Flags"
        value="4"
        description="Across all environments"
      />
      <StatCard title="Total Environments" value="3" description="registered" />
      <StatCard
        title="Total Evaluations"
        value="10"
        description="Last 24 hours"
      />
      <StatCard
        title="Evaluation Time"
        value="2ms"
        description="Average latency when evaluating"
      />
    </div>
  );
}
