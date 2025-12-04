import { PeopleAltOutlined, ViewAgendaOutlined } from '@mui/icons-material';
import { generatePath, matchPath, useLocation, useParams } from "react-router-dom";
import { absoluteRouteMap } from "@agent-management-platform/types";
import type { NavigationItem, NavigationSection } from '@agent-management-platform/views';
// import { useGetAgent } from '@agent-management-platform/api-client';

export function useNavigationItems(): Array<NavigationSection | NavigationItem> {
  const { orgId, projectId } = useParams();
  const { pathname } = useLocation();
  if (orgId && projectId) {
    return [
      {
        label: 'Agents',
        type: 'item',
        icon: <PeopleAltOutlined fontSize='small' />,
        href: generatePath(
          absoluteRouteMap.children.org.children.projects.path,
          { orgId, projectId }),
        isActive: !!matchPath(absoluteRouteMap.children.org.children.projects.path, pathname) ||
          !!matchPath(absoluteRouteMap.
            children.org.children.projects.children.agents.wildPath, pathname),
      },
    ]
  }
  if (orgId) {
    return [
      {
        label: 'Projects',
        type: 'item',
        icon: <ViewAgendaOutlined fontSize='small' />,
        href: generatePath(
          absoluteRouteMap.children.org.path,
          { orgId }),
        isActive: !!matchPath(absoluteRouteMap.children.org.path, pathname),
      },
    ]
  }
  return []
}
