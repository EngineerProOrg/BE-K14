import { Box, Container, Grid, Paper } from "@mui/material";
import Header from "./common/Header";
import Footer from "./common/Footer";
import LeftSidebar from "./common/LeftSidebar";
import RightSidebar from "./common/RightSidebar";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <Box display="flex" flexDirection="column" minHeight="100vh">
      <Header />
      <Container maxWidth="lg" sx={{ flexGrow: 1, mt: 2 }}>
        <Grid container spacing={2}>
          <Grid item xs={12} md={3} component="div">
            <Paper elevation={2} sx={{ p: 2 }}>
              <LeftSidebar />
            </Paper>
          </Grid>
          <Grid item xs={12} md={6} component="div">
            {children}
          </Grid>
          <Grid item xs={12} md={3} component="div">
            <Paper elevation={2} sx={{ p: 2 }}>
              <RightSidebar />
            </Paper>
          </Grid>
        </Grid>
      </Container>
      <Footer />
    </Box>
  );
}
