import React from 'react';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './pages/Home';
import Design from './pages/Design';
import {QueryClient, QueryClientProvider} from 'react-query';
import {Helmet} from 'react-helmet';

const queryClient = new QueryClient();

const App: React.FC = () => {
    return (
        <QueryClientProvider client={queryClient}>
            <Helmet>
                <body className="max-w-5xl mx-auto"/>
            </Helmet>
            <Router>
                <Navbar/>
                <Switch>
                    <Route exact path="/site/:id" children={<Home/>}/>
                    <Route exact path="/design">
                        <Design/>
                    </Route>
                </Switch>
            </Router>
        </QueryClientProvider>
    );
};

export default App;
