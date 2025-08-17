import { Navigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export default function PrivateRoute({ children, requiredRoles }) {
  const { isAuthenticated, userRole } = useAuth();

  if (!isAuthenticated) return <Navigate to="/" />;

  if (requiredRoles && !requiredRoles.includes(userRole)) return <Navigate to="/" replace />;

  return children;
}
