
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Layout } from "../Layouts";
import { Protected } from "../Providers/Protected";
import {
  addNewAgentPageMetaData,
  overviewMetadata,
  deployMetadata,
  testMetadata,
  tracesMetadata,
  Login
} from "../pages";
import { relativeRouteMap } from "@agent-management-platform/types";
import { AgentNavBar, EnvSubNavBar } from "@agent-management-platform/shared-component";
export function RootRouter() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path={relativeRouteMap.children.login.path} element={<Login />} />
        <Route
          path={"/"}
          element={<Protected><Layout /></Protected>}>
          <Route path={relativeRouteMap.children.org.path}>
            <Route
              index
              element={<overviewMetadata.levels.organization />}
            />
            <Route
              path={relativeRouteMap.children.org.children.projects.path}
            >
              <Route
                index
                element={<overviewMetadata.levels.project />}
              />
              <Route
                path={relativeRouteMap.children.org.children.
                  projects.children.newAgent.path}
                element={<addNewAgentPageMetaData.component />}
              />
              <Route
                path={relativeRouteMap.children.org.children.
                  projects.children.agents.path}
                element={<AgentNavBar />}
              >
                <Route
                  index
                  element={<overviewMetadata.levels.component />}
                />
                <Route
                  path={relativeRouteMap.children.org.children.
                    projects.children.agents.children.traces.path}
                  element={<overviewMetadata.levels.component />}
                />
                <Route
                  path={relativeRouteMap.children.org.children.
                    projects.children.agents.children.environment.path}
                  element={<EnvSubNavBar />}
                >
                  <Route
                    index
                    element={<deployMetadata.levels.component />}
                  />
                  <Route
                    path={relativeRouteMap.children.org.children.
                      projects.children.agents.children.environment.children.tryOut.path}
                    element={<testMetadata.levels.component />}
                  />
                  <Route
                    path={relativeRouteMap.children.org.children.
                      projects.children.agents.children.environment.children.observability.path + "/*"}
                    element={<tracesMetadata.levels.component />}
                  />
                </Route>
              </Route>
            </Route>
            <Route path="*" element={<div>404 Not Found</div>} />
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}
