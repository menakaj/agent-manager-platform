import { Box, IconButton, Tooltip } from "@mui/material";
import { ContentCopy } from "@mui/icons-material";
import { useState } from "react";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { oneDark } from "react-syntax-highlighter/dist/esm/styles/prism";

interface CodeBlockProps {
    code: string;
    language?: string;
    showCopyButton?: boolean;
    fieldId?: string;
}

export const CodeBlock = ({ 
    code, 
    language = "bash", 
    showCopyButton = true,
    fieldId = "code" 
}: CodeBlockProps) => {
    const [copiedField, setCopiedField] = useState<string | null>(null);

    const handleCopy = async () => {
        try {
            await navigator.clipboard.writeText(code);
            setCopiedField(fieldId);
            setTimeout(() => setCopiedField(null), 2000);
        } catch {
            // Failed to copy - silently fail
        }
    };

    return (
        <Box
            sx={{
                position: 'relative',
                borderRadius: 1,
                overflow: 'hidden',
                '& pre': {
                    margin: 0,
                    padding: 2,
                    borderRadius: 1,
                    fontSize: '0.875rem',
                }
            }}
        >
            {showCopyButton && (
                <Tooltip title={copiedField === fieldId ? 'Copied!' : 'Copy code'}>
                    <IconButton
                        onClick={handleCopy}
                        size="small"
                        sx={{
                            position: 'absolute',
                            right: 8,
                            top: 8,
                            zIndex: 1,
                            color: 'grey.400',
                            bgcolor: 'rgba(0, 0, 0, 0.3)',
                            '&:hover': {
                                bgcolor: 'rgba(0, 0, 0, 0.5)',
                            }
                        }}
                    >
                        <ContentCopy fontSize="small" />
                    </IconButton>
                </Tooltip>
            )}
            <SyntaxHighlighter
                language={language}
                style={oneDark}
                customStyle={{
                    margin: 0,
                    borderRadius: '4px',
                }}
            >
                {code}
            </SyntaxHighlighter>
        </Box>
    );
};

