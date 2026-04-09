import React from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import {
  MessageSquare,
  LayoutDashboard,
  LogOut,
  Settings,
  Bell,
  Search,
  User
} from 'lucide-react';

const Layout = ({ children }) => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = () => {
    logout();
    navigate('/login');
  };

  return (
    <div className="app-container">
      {/* Sidebar */}
      <div className="sidebar bg-bg-card p-4">
        <div className="p-6 mb-4 flex items-center gap-3">
          <span style={{ fontSize: '20px', fontWeight: '800', letterSpacing: '-0.5px' }}>Sociomile</span>
        </div>

        <div className="p-4 border-t" style={{ borderColor: 'var(--border)' }}>
          <div className="glass p-4 mb-4 flex items-center gap-3">
            <div className="w-10 h-10 rounded-full bg-primary/20 flex items-center justify-center">
              <User size={20} className="text-primary" />
            </div>
            <div className="overflow-hidden">
              <p style={{ fontSize: '14px', fontWeight: '600' }} className="truncate">
                {user?.email.split('@')[0]}
              </p>
              <p style={{ fontSize: '11px', color: 'var(--text-secondary)' }} className="truncate">
                {user?.email}
              </p>
            </div>
          </div>
          <button
            onClick={handleLogout}
            className="w-full flex items-center gap-3 px-4 py-3 text-error hover:bg-error/5 rounded-xl transition-all"
            style={{ textAlign: 'left' }}
          >
            <LogOut size={20} />
            <span style={{ fontWeight: '600' }}>Sign Out</span>
          </button>
        </div>
      </div>

      {/* Main Content */}
      <div className="main-content">
        {/* Page Content */}
        <main className="flex-1 overflow-y-auto p-4 space-y-6">
          {children}
        </main>
      </div>
    </div>
  );
};

export default Layout;
