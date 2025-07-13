import "./index.css";

import { HeadRow } from "./components/HeadRow";
import AppHeader from "./components/Header";
import FeatureList from "./components/FeatureList";
import { SidebarProvider } from "./components/ui/sidebar";
import { AppSidebar } from "./components/AppSidebar";
export function App() {
  return (
    <SidebarProvider>
      <div className="container mx-auto p-8 text-center relative z-10">
        <AppSidebar />
        <AppHeader />
        <HeadRow />
        <FeatureList />
      </div>
    </SidebarProvider>
  );
}

export default App;
