import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

interface StatCardProps {
  title: string;
  description: string;
  value: string;
}

export default function StatCard({ title, description, value }: StatCardProps) {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">{title}</CardTitle>
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold text-left">{value}</div>
        <p className="text-xs text-muted-foreground text-left">{description}</p>
      </CardContent>
    </Card>
  );
}
