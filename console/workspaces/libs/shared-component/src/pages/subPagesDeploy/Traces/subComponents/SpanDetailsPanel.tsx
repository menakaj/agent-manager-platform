import { Box, Divider, useTheme } from "@mui/material";
import { Span } from "@agent-management-platform/types";
import { SpanPanelHeader } from "./spanDetails/SpanPanelHeader";
import { BasicInfoSection } from "./spanDetails/BasicInfoSection";
import { TimingSection } from "./spanDetails/TimingSection";
import { StatusSection } from "./spanDetails/StatusSection";
import { AttributesSection } from "./spanDetails/AttributesSection";

interface SpanDetailsPanelProps {
    span: Span | null;
    onClose: () => void;
}

export function SpanDetailsPanel({ span, onClose }: SpanDetailsPanelProps) {
    const theme = useTheme();

    if (!span) {
        return null;
    }

    return (
        <Box
            sx={{
                width: theme.spacing(80),
                p: theme.spacing(2),
                height: '100%',
                display: 'flex',
                flexDirection: 'column',
                gap: theme.spacing(2),
                bgcolor: theme.palette.background.paper
            }}
        >
            <SpanPanelHeader onClose={onClose} />
            <Divider />
            
            <Box 
                sx={{ 
                    display: 'flex', 
                    flexDirection: 'column', 
                    gap: theme.spacing(2), 
                    overflow: 'auto', 
                    flex: 1 
                }}
            >
                <BasicInfoSection span={span} />
                <Divider />
                <TimingSection span={span} />
                <Divider />
                <StatusSection span={span} />
                {span.attributes && Object.keys(span.attributes).length > 0 && (
                    <>
                        <Divider />
                        <AttributesSection attributes={span.attributes} />
                    </>
                )}
            </Box>
        </Box>
    );
}

