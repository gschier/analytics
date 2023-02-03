import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Navbar from './components/Navbar';
import Site from './pages/Site';
import Design from './pages/Design';
import Home from './pages/Home';
import { QueryClient, QueryClientProvider } from 'react-query';
import { Helmet } from 'react-helmet';
import useStateLocalStorage from './hooks/use-state-localstorage';
import { useTheme } from './hooks/use-theme';

const queryClient = new QueryClient();

const App: React.FC = () => {
  const [theme] = useTheme();
  return (
    <QueryClientProvider client={queryClient}>
      <Helmet>
        <body className="max-w-5xl mx-auto" />
        <html className={theme} />
      </Helmet>
      <Router>
        <Navbar />
        <Switch>
          <Route exact path="/">
            <Home />
          </Route>
          <Route exact path="/analytics/:id" children={<Site />} />
          <Route exact path="/design">
            <Design />
          </Route>
        </Switch>
      </Router>
    </QueryClientProvider>
  );
};

export default App;
