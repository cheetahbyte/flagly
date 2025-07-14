export type Flag = {
  key: string;
  description: string;
  environments: Record<string, Environment>;
  usage: number;
};

export type Rollout = {
  percentage?: number;
  stickiness?: number;
};

export type Environment = {
  enabled: boolean;
  rollout?: Rollout;
};

export type ApiStatusResponse = {
  version: string;
};
