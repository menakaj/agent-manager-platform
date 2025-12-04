/* eslint-disable no-console */

import { rootRouteMap } from "../src/routes/routes.map";
import { AppRoute, GenaratedRoute } from "@agent-management-platform/types";
import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);


const traverseOnRoute = (route: AppRoute, parentPath?: string) => {
    const routePath = parentPath ? `${parentPath}/${route.path}` : route.path;
    const wildPath = routePath ? `${routePath}/*` : '*';
    const generatedRoute: GenaratedRoute = {
        path: routePath,
        wildPath: wildPath,
        children: {},
    };
    if (route.children) {
        Object.entries(route.children).forEach(([key, child]) =>{
            generatedRoute.children = {
                ...generatedRoute.children, 
                [key]: traverseOnRoute(child, routePath)
            };
        });
    }
    return generatedRoute;
}

export const genarateRoutes = () => {
    return  traverseOnRoute(rootRouteMap);
}


const routes = JSON.stringify(genarateRoutes(), null, 2);
const outputPath = path.resolve(__dirname, '../src/routes/generated-route.map.ts');
fs.writeFileSync(outputPath, `export const generatedRouteMap =  ${routes};`);

console.log('✅ Route map generated successfully!');
console.log('📁 Output file: src/routes/generated-route.map.ts');
console.log('🔄 Run this command again to regenerate the routes');
