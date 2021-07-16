import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Navbar from './molecules/Navbar';
import Home from './pages/Home';
import Design from './pages/Design';

const App: React.FC = () => {
    return (
        <Router>
            <Navbar />
            <Switch>
                <Route exact path="/">
                    <Home />
                </Route>
                <Route exact path="/design">
                    <Design />
                </Route>
            </Switch>
        </Router>
    );
};

export default App;
