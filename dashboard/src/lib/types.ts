export type Flag = {
  key: string;
  description: string;
  environments: Record<string, Environment>;
  usage: number;
};

export type Environment = {
  enabled: boolean;
  rollout?: any;
};

export type ApiStatusResponse = {
  version: string;
};
