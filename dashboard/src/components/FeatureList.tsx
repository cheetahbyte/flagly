import React, { useState } from "react";
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

// Utility function to convert snake_case_key to Title Case
const formatKeyAsTitle = (key) => {
  return key
    .split("_") // Split by underscore
    .map((word) => word.charAt(0).toUpperCase() + word.slice(1)) // Capitalize first letter of each word
    .join(" "); // Join with space
};

export default function FeatureList() {
  // Global state for the selected environment
  const [selectedGlobalEnvironment, setSelectedGlobalEnvironment] =
    useState("production");

  // Updated feature data structure
  const [features, setFeatures] = useState([
    {
      key: "new_dashboard_ui",
      description: "Enable the redesigned dashboard interface with improved UX",
      environments: {
        production: {
          enabled: true,
          rollout: { percentage: 85, stickiness: "user_id" }, // Has explicit rollout
        },
        staging: { enabled: true }, // Should implicitly be 100%
        development: { enabled: false }, // Should implicitly be 0%
      },
      usage: "85%", // This will be the main displayed usage
      // evaluations: "12,450", // Removed from data as per requirement
    },
    {
      key: "advanced_analytics",
      description: "Show advanced analytics features to premium users",
      environments: {
        production: {
          enabled: false,
          rollout: { percentage: 0, stickiness: "user_id" }, // Has explicit rollout of 0
        },
        staging: { enabled: false }, // Should implicitly be 0%
        development: { enabled: true }, // Should implicitly be 100%
      },
      usage: "0%",
      // evaluations: "8,920", // Removed from data as per requirement
    },
    {
      key: "beta_features_enabled",
      description: "Enable access to beta features for selected users",
      environments: {
        production: {
          enabled: false,
          rollout: { percentage: 0, stickiness: "user_id" }, // Has explicit rollout of 0
        },
        staging: { enabled: true }, // Should implicitly be 100%
        development: { enabled: false }, // Should implicitly be 0%
      },
      usage: "0%",
      // evaluations: "5,670", // Removed from data as per requirement
    },
    {
      key: "search_bar_improvements",
      description: "Improvements to the main search bar functionality",
      environments: {
        production: { enabled: false, rollout: { percentage: 10 } }, // Has explicit rollout
        staging: { enabled: true },
        development: { enabled: true },
      },
      usage: "15%",
      // evaluations: "2,100", // Removed from data as per requirement
    },
  ]);

  // Extract all unique environment keys from the features data
  const allEnvironments = Array.from(
    new Set(features.flatMap((f) => Object.keys(f.environments)))
  );

  return (
    <div className="space-y-4 p-4">
      {/* Global Environment Dropdown at the top */}
      <div className="mb-6 flex items-center justify-between gap-4">
        <Alert className={`flex w-fit my-custom-alert-class`}>
          <InfoIcon />
          <AlertTitle className="text-left">
            System is running in immutable mode.
          </AlertTitle>
        </Alert>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="my-custom-button-class">
              Environment:{" "}
              {selectedGlobalEnvironment.charAt(0).toUpperCase() +
                selectedGlobalEnvironment.slice(1)}
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

      {/* Feature List */}
      {features.map((feature) => {
        const environmentData = feature.environments[selectedGlobalEnvironment];
        const isFeatureEnabled = environmentData?.enabled || false;

        // Calculate the effective rollout percentage
        let effectiveRolloutPercentage;
        if (environmentData?.rollout?.percentage !== undefined) {
          effectiveRolloutPercentage = environmentData.rollout.percentage;
        } else {
          // If no explicit rollout percentage, infer from 'enabled' state
          effectiveRolloutPercentage = isFeatureEnabled ? 100 : 0;
        }

        return (
          <Card key={feature.key}>
            <CardContent className="flex items-center justify-between p-6">
              <div className="flex items-center space-x-4">
                {/* Custom styling for red/green switch */}
                <Switch
                  id={feature.key}
                  checked={isFeatureEnabled}
                  // Make the switch non-interactive
                  disabled={true} // Disables interaction
                  aria-readonly="true" // Semantic non-interactivity
                  // Custom classes for red/green background based on checked state
                  className={
                    isFeatureEnabled
                      ? "bg-green-500 data-[state=checked]:bg-green-500"
                      : "bg-red-500 data-[state=unchecked]:bg-red-500"
                  }
                />
                <div>
                  <h3 className="text-lg font-semibold text-gray-900 text-left">
                    {formatKeyAsTitle(feature.key)}
                  </h3>
                  <p className="text-sm text-gray-600">{feature.description}</p>
                </div>
              </div>
              <div className="flex flex-col items-end text-right ml-auto">
                {/* Display Usage (bigger) */}
                <div className="text-xl font-bold text-gray-900">
                  {feature.usage} usage
                </div>
                {/* Display Rollout Percentage (smaller, underneath) */}
                <div className="text-xs text-muted-foreground">
                  {effectiveRolloutPercentage}% rollout
                </div>
              </div>
            </CardContent>
          </Card>
        );
      })}
    </div>
  );
}
