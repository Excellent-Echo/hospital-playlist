import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'
import PrivateRouteAdmin from './PrivateRouteAdmin'

//
import NavbarMenu from '../components/NavbarMenu'

//
import Homepage from '../pages/Homepage'
import NotFound from '../pages/NotFound'
import Login from '../pages/Login'
import Register from '../pages/Register'
import Profile from '../pages/Profile'
import AdminDashboard from '../pages/Admin/Dashboard'

//

const Routes = () => {
	let userRole = localStorage.getItem('userRole')
	let getIsAuth = localStorage.getItem('isAuth')
	let isAuth = JSON.parse(getIsAuth)
	return (
		<>
			<Router>
				<NavbarMenu />
				<Switch>
					<Route path="/register" exact component={Register} />
					<PrivateRoute path="/profile" component={Profile} auth={isAuth} role={userRole} />
					<PrivateRouteAdmin path="/admin" component={AdminDashboard} auth={isAuth} role={userRole} />
					<Route path="/login" exact>
						<Login />
					</Route>
					<Route path="/" exact component={Homepage} />
					<Route component={NotFound} />
				</Switch>
			</Router>
		</>
	)
}

export default Routes
