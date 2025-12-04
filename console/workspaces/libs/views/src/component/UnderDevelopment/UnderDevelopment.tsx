import { Box, Typography, useTheme } from "@mui/material";
import { FadeIn } from "../FadeIn/FadeIn";
import Image from "../Image/Image";
import { ImageList } from "../Image";

export function UnderDevelopment() {
    const theme = useTheme();
    return (
        <FadeIn>
            <Box sx={{
                display: 'flex',
                flexDirection: 'column',
                justifyContent: 'center',
                alignItems: 'center',
                height: '70vh',
                p: theme.spacing(2),
                gridGap: theme.spacing(2)
            }}>
                <Image
                    src={ImageList.UNDER_DEVELOPMENT} alt="Under Development"
                    width={"30%"}
                    height={"30%"}
                />
                <Typography variant="h5" align="center" color="secondary.main">
                    Coming Soon...
                </Typography>
                <Typography variant="body1" align="center" color="text.secondary">
                    This functionality is currently under development
                    and will be released in an upcoming release.
                </Typography>
            </Box>
        </FadeIn>
    );
}

