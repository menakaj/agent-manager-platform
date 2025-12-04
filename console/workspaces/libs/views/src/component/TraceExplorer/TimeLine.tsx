import { Box, Typography, useTheme } from "@mui/material";
import { useMemo } from "react";
import dayjs from "dayjs";

interface TimeLineProps {
    startTime: string;
    endTime: string;
    parentStartTime?: string;
    parentEndTime?: string;
}
export function TimeLine(props: TimeLineProps) {
    const { startTime, endTime, parentStartTime, parentEndTime } = props;
    const theme = useTheme();
    const { width, startPosition } = useMemo(() => {
        if (!parentStartTime && !parentEndTime) {
            return {
                width: '100%',
                startPosition: '0%'
            };
        }
        const startTimeMs = new Date(startTime).getTime();
        const endTimeMs = new Date(endTime).getTime();
        const parentStartTimeMs = parentStartTime ? new Date(parentStartTime).getTime() : 0;
        const parentEndTimeMs = parentEndTime ? new Date(parentEndTime).getTime() : 0;
        return {
            width: `${((endTimeMs - startTimeMs) * 100 / (parentEndTimeMs - parentStartTimeMs)).toFixed(2)}%`,
            startPosition: (100 * (startTimeMs - parentStartTimeMs) / (parentEndTimeMs - parentStartTimeMs)).toFixed(2) + '%'
        };
    }, [startTime, endTime, parentStartTime, parentEndTime]);
    return (
        <Box
            position="relative"
            width={300}
            display="flex"
            flexDirection="column"
            gap={1}
            pt={0.25}
        >
            <Box
                height={theme.spacing(0.5)}
                bgcolor={theme.palette.text.disabled}
                width={"100%"}
                borderRadius={theme.spacing(0.5)}
            />
            <Box
                position="absolute"
                borderRadius={theme.spacing(0.5)}
                top={0}
                left={startPosition}
                width={width}
                sx={{
                    background: `linear-gradient(45deg, ${theme.palette.secondary.main}, ${theme.palette.primary.main})`,
                }}
                height={theme.spacing(1)}
            />
            <Typography variant="caption">
                Executed from {dayjs(startTime).format('HH:mm:ss.SSS')}
                &nbsp;to&nbsp;
                {dayjs(endTime).format('HH:mm:ss.SSS')}
            </Typography>
        </Box>
    );
}
