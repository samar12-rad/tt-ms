import React from 'react';
import { useLocation } from 'react-router-dom';
import NavBar from './NavBar';
import { useAuth } from '../context/AuthContext';

const Layout = ({ children }) => {
  const location = useLocation();
  const { logout } = useAuth();
  
  const noNavBarRoutes = ['/'];
  
  const shouldShowNavBar = !noNavBarRoutes.includes(location.pathname);

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 via-blue-50 to-indigo-50">
      {/* Conditionally render NavBar */}
      {shouldShowNavBar && <NavBar onLogout={logout} />}
      <main className={shouldShowNavBar ? 'pt-0' : ''}>
        {children}
      </main>
    </div>
  );
};

export default Layout;