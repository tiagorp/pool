import './App.css';
import { Route, Routes } from 'react-router-dom';
import Layout from './components/MainLayout/MainLayout';
import Home from './routes/Home';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import PageLayout from './components/PageLayout/PageLayout';
import About from './routes/About';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />

      <Routes>
        <Route path="/" element={<Layout />}>
          <Route
            index
            element={
              <PageLayout>
                <Home />
              </PageLayout>
            }
          />

          <Route
            path="about"
            element={
              <PageLayout>
                <About />
              </PageLayout>
            }
          />

          {/* Using path="*"" means "match anything", so this route
            acts like a catch-all for URLs that we don't have explicit
            routes for. */}
          {/* <Route path="*" element={<NoMatch />} /> */}
        </Route>
      </Routes>
    </ThemeProvider>
  );
}

export default App;
