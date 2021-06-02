import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import PrivateRoute from './PrivateRoute'
import { useSelector } from 'react-redux'

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
	let isAuth = localStorage.getItem('isAuth')
	console.log(typeof isAuth)
	let x = JSON.parse(isAuth)
	console.log(typeof x)
	// const userLoginData = useSelector((state) => state.userLogin)
	return (
		<>
			<Router>
				<NavbarMenu />
				<Switch>
					<Route path="/register" exact component={Register} />
					<PrivateRoute path="/profile" component={Profile} auth={x} />
					<PrivateRoute path="/admin" component={AdminDashboard} auth={x} />
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
