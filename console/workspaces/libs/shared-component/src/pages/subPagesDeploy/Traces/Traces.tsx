import React, { useState } from 'react';
import { Box, Collapse, useTheme } from '@mui/material';
import { TopCards } from './subComponents/TopCards';
import { TracesTable } from './subComponents/TracesTable';
import { FadeIn } from '@agent-management-platform/views';
import { Route, Routes, useMatch } from 'react-router-dom';
import { TraceDetails } from './subComponents/TraceDetails';
import { TraceListTimeRange } from '@agent-management-platform/types';

export const Traces: React.FC = () => {
  const theme = useTheme();
  const [timeRange, setTimeRange] = useState<TraceListTimeRange>(TraceListTimeRange.ONE_DAY);
  const isTraceDetails = useMatch("/unknown");

  return (
    <FadeIn>
      <Box 
        sx={{ 
          pt: theme.spacing(1), 
          gap: theme.spacing(2), 
          display: 'flex', 
          flexDirection: 'column' 
        }}
      >
        <Collapse in={!isTraceDetails}>
          <TopCards timeRange={timeRange} />
        </Collapse>
        <Routes>
          <Route path="/" element={<TracesTable timeRange={timeRange} setTimeRange={setTimeRange} />} />
          <Route path="/:traceId" element={<TraceDetails />} />
        </Routes>
      </Box>
    </FadeIn>
  );
};

