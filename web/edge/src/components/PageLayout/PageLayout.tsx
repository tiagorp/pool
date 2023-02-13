import { Box } from "@mui/system";
import { ReactNode } from "react";


export default function PageLayout(props: React.PropsWithChildren) {
    return (
        <Box sx={{p: 4}}>
        {props.children}
        </Box>
    )
}