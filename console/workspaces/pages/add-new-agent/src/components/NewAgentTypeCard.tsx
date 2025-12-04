import { Card, CardActionArea, CardContent, Box, Typography, alpha } from "@mui/material";
import { ArrowForward } from "@mui/icons-material";

interface NewAgentTypeCardProps {
    type: string;
    title: string;
    subheader: string;
    icon: React.ReactNode;
    content: React.ReactNode;
    onClick: (type: string) => void;
}

export const NewAgentTypeCard = (props: NewAgentTypeCardProps) => {
    const { type, title, subheader, icon, content, onClick } = props;

    const handleClick = () => {
        onClick(type);
    };

    return (
        <Card
            variant="outlined"
            sx={{
                height: '100%',
                transition: 'all 0.3s ease-in-out',
                '&:hover': {
                    boxShadow: 4,
                    borderColor: 'primary.main',
                    transform: 'translateY(-4px)',
                },
            }}
        >
            <CardActionArea
                onClick={handleClick}
                sx={{
                    height: '100%',
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'stretch',
                    justifyContent: 'flex-start',
                }}
            >
                <CardContent sx={{ flexGrow: 1, width: '100%', p: 3 }}>
                    {/* Icon Header */}
                    <Box
                        sx={{
                            mb: 2,
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center',
                            width: 64,
                            height: 64,
                            borderRadius: 2,
                            bgcolor: (theme) => alpha(theme.palette.primary.main, 0.1),
                            color: 'primary.main',
                        }}
                    >
                        {icon}
                    </Box>

                    {/* Title and Subheader */}
                    <Typography variant="h5" component="div" gutterBottom>
                        {title}
                    </Typography>
                    <Typography variant="body2" color="text.secondary" sx={{ mb: 2 }}>
                        {subheader}
                    </Typography>

                    {/* Content */}
                    <Box sx={{ mb: 2 }}>
                        {content}
                    </Box>

                    {/* Call to Action */}
                    <Box
                        sx={{
                            display: 'flex',
                            alignItems: 'center',
                            gap: 1,
                            color: 'primary.main',
                            mt: 'auto',
                        }}
                    >
                        <Typography variant="body2" fontWeight="medium">
                            Get Started
                        </Typography>
                        <ArrowForward fontSize="small" />
                    </Box>
                </CardContent>
            </CardActionArea>
        </Card>
    );
};
