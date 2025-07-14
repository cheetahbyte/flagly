import React, { useEffect, useState } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { Switch } from "@/components/ui/switch";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuRadioGroup,
  DropdownMenuRadioItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";
import { InfoIcon } from "lucide-react";
import type { Environment, Flag } from "@/lib/types";
import { Badge } from "./ui/badge";

interface FeatureListProps {
  flags: Flag[];
  environments: string[];
}

const formatKeyAsTitle = (key) => {
  return key
    .split("_") // Split by underscore
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1)) // Capitalize first letter of each word
    .join(" "); // Join with space
};

export default function FeatureList({ flags, environments }: FeatureListProps) {
  const [selectedGlobalEnvironment, setSelectedGlobalEnvironment] =
    useState("");

  const allEnvironments = Array.from(
    new Set(flags.flatMap((f) => Object.keys(f.environments)))
  );

  useEffect(() => {
    setSelectedGlobalEnvironment(environments[0]);
  }, [environments]);

  return (
    <div className="space-y-4 p-4">
      <div className="mb-6 flex items-center justify-between gap-4">
        <Alert className={`flex w-fit my-custom-alert-class`}>
          <InfoIcon />
          <AlertTitle className="text-left">
            System is running in immutable mode.{" "}
            <a href="https://github.com/cheetahbyte/flagly/wiki/Immutable-Mode--&-GitOps-Mode">
              Learn more
            </a>
          </AlertTitle>
        </Alert>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="my-custom-button-class">
              Environment:{" "}
              {selectedGlobalEnvironment
                ? selectedGlobalEnvironment.charAt(0).toUpperCase() +
                  selectedGlobalEnvironment.slice(1)
                : "N/A"}
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-56">
            <DropdownMenuLabel>Select Environment</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuRadioGroup
              value={selectedGlobalEnvironment}
              onValueChange={setSelectedGlobalEnvironment}
            >
              {allEnvironments.map((envKey) => (
                <DropdownMenuRadioItem key={envKey} value={envKey}>
                  {envKey.charAt(0).toUpperCase() + envKey.slice(1)}
                </DropdownMenuRadioItem>
              ))}
            </DropdownMenuRadioGroup>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      {flags.map((feature) => {
        const environmentData = feature.environments[selectedGlobalEnvironment];
        const isFeatureEnabled = environmentData?.enabled || false;
        let effectiveRolloutPercentage = 0; // Default to 0
        const rolloutPercentage = environmentData?.rollout?.percentage;
        if (environmentData?.enabled && rolloutPercentage === 0) {
          effectiveRolloutPercentage = 100;
        } else if (rolloutPercentage !== undefined) {
          effectiveRolloutPercentage = rolloutPercentage;
        }

        return (
          <Card key={feature.key}>
            <CardContent className="flex items-center justify-between p-6">
              <div className="flex items-center space-x-4">
                <Switch
                  id={feature.key}
                  checked={isFeatureEnabled}
                  disabled={true}
                  aria-readonly="true"
                  className={
                    isFeatureEnabled
                      ? "bg-green-500 data-[state=checked]:bg-green-500"
                      : "bg-red-500 data-[state=unchecked]:bg-red-500"
                  }
                />
                <div>
                  <h3 className="text-lg font-semibold text-gray-900 text-left flex items-center gap-2">
                    {formatKeyAsTitle(feature.key)}{" "}
                    <Badge variant={"outline"}>{feature.key}</Badge>
                  </h3>
                  <p className="text-sm text-gray-600 text-left">
                    {feature.description}
                  </p>
                </div>
              </div>
              <div className="flex flex-col items-end text-right ml-auto">
                {/* TODO: enable this again later */}
                {/* <div className="text-xl font-bold text-gray-900">
                  {feature.usage} usage
                </div>
                <div className="text-xs text-muted-foreground">
                  {effectiveRolloutPercentage}% rollout
                </div> */}
                <div className="text-xl font-bold text-gray-900">
                  {environmentData?.enabled &&
                    effectiveRolloutPercentage + "%rollout"}
                </div>
              </div>
            </CardContent>
          </Card>
        );
      })}
    </div>
  );
}
