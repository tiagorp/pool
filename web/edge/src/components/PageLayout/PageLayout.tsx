import { Box } from "@mui/system";

export default function PageLayout(props: React.PropsWithChildren) {
  return <Box sx={{ p: 4 }}>{props.children}</Box>;
}
