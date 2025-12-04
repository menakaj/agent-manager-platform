import { AgentsListPage } from './AgentsListPage';
import {Person} from "@mui/icons-material";
import { absoluteRouteMap } from '@agent-management-platform/types';

export const metaData = {
  title: 'Agents',
  description: 'Agents List Page',
  icon: Person,
  path: absoluteRouteMap.children.org.children.projects.path,
  component: AgentsListPage,
}
