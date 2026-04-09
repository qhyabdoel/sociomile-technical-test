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

  const navItems = [
    { icon: LayoutDashboard, label: 'Dashboard', path: '/' },
    { icon: MessageSquare, label: 'Conversations', path: '/conversations' },
    { icon: Bell, label: 'Notifications', path: '/notifications' },
    { icon: Settings, label: 'Settings', path: '/settings' },
  ];

  return (
    <div className="app-container">
      {/* Sidebar */}
      <div className="sidebar bg-bg-card">
        <div className="p-6 mb-4 flex items-center gap-3">
          <div className="w-10 h-10 bg-primary rounded-xl flex items-center justify-center">
            <MessageSquare color="white" size={24} />
          </div>
          <span style={{ fontSize: '20px', fontWeight: '800', letterSpacing: '-0.5px' }}>Sociomile</span>
        </div>

        <nav className="flex-1 px-4 space-y-2">
          {navItems.map((item) => (
            <button
              key={item.label}
              onClick={() => navigate(item.path)}
              className={`w-full flex items-center gap-3 px-4 py-3 rounded-xl transition-all ${
                location.pathname === item.path 
                  ? 'bg-primary text-white' 
                  : 'text-text-secondary hover:bg-white/5'
              }`}
              style={{ textAlign: 'left' }}
            >
              <item.icon size={20} />
              <span style={{ fontWeight: '500' }}>{item.label}</span>
            </button>
          ))}
        </nav>

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
        {/* Top Header */}
        <header className="h-16 px-8 border-b flex items-center justify-between" style={{ borderColor: 'var(--border)' }}>
          <div className="flex-1 max-w-xl">
            <div className="relative">
              <Search className="absolute left-3 top-2.5 text-text-secondary" size={18} />
              <input 
                placeholder="Search conversations..." 
                className="pl-10 h-10 text-sm"
                style={{ backgroundColor: 'rgba(255,255,255,0.03)' }}
              />
            </div>
          </div>
          <div className="flex items-center gap-4">
            <div className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center cursor-pointer hover:bg-white/10">
              <Bell size={18} className="text-text-secondary" />
            </div>
            <div className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center cursor-pointer hover:bg-white/10">
              <Settings size={18} className="text-text-secondary" />
            </div>
          </div>
        </header>

        {/* Page Content */}
        <main className="flex-1 overflow-y-auto">
          {children}
        </main>
      </div>
    </div>
  );
};

export default Layout;
