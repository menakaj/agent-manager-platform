import { type AppRoute } from "./types";

export const rootRouteMap: AppRoute = {
    path: '',
    children: {
    login: {
        path: '/login',
        index: true,
        children: {},
    },
    org: {
        path: '/org/:orgId',
        index: true,
        children: {
            newProject: {
                path: 'newProject',
                index: true,
                children: {},
            },
            projects: {
                path: 'project/:projectId',
                index: true,
                children: {
                    newAgent: {
                        path: 'newAgent',
                        index: true,
                        children: {},
                    },
                    agents: {
                        path: 'agents/:agentId',
                        index: true,
                        children: {
                            traces:{
                                // Note that this is for external component traces
                                path: 'traces/:traceId',
                                index: true,
                                children: {},
                            },
                            build: {
                                path: 'build',
                                index: true,
                                children: {},
                            },
                            environment:{
                                path: "environment/:envId",
                                index:false,
                                children:{
                                    deploy: {
                                        path: 'deploy',
                                        index: true,
                                        children: {},
                                    },
                                    tryOut: {
                                        path: 'tryOut',
                                        index: true,
                                        children: {},
                                    },
                                    observability: {
                                        path: 'observability',
                                        index: true,
                                        children: {
                                            traces: {
                                                path: 'traces/:traceId',
                                                index: true,
                                                children: {},
                                            },
                                        },
                                    },
                                }
                            },
                        },
                    },
                },
            },
        },
    },
    },
}
