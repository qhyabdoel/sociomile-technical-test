-- create tenants if not exist
CREATE TABLE IF NOT EXISTS tenants (
    id BiGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create users if not exist
CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role ENUM('admin', 'agent') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    INDEX idx_tenant_id (tenant_id)
);

-- create conversations if not exist
CREATE TABLE IF NOT EXISTS conversations (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    customer_external_id VARCHAR(255) NOT NULL,
    status ENUM('open', 'closed') NOT NULL DEFAULT 'open',
    assigned_agent_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    FOREIGN KEY (assigned_agent_id) REFERENCES users(id)
);

-- create messages if not exist
CREATE TABLE IF NOT EXISTS messages (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    conversation_id BIGINT NOT NULL,
    sender_type ENUM('customer', 'agent') NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (conversation_id) REFERENCES conversations(id),
    INDEX idx_conversation_id (conversation_id)
);

-- create tickets if not exist
CREATE TABLE IF NOT EXISTS tickets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    conversation_id BIGINT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status ENUM('open', 'in_progress', 'resolved', 'closed') NOT NULL DEFAULT 'open',
    priority ENUM('low', 'medium', 'high', 'urgent') NOT NULL DEFAULT 'medium',
    assigned_agent_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (tenant_id) REFERENCES tenants(id),
    FOREIGN KEY (conversation_id) REFERENCES conversations(id),
    FOREIGN KEY (assigned_agent_id) REFERENCES users(id),
    INDEX idx_tenant_id (tenant_id)
);

-- seed tenants
INSERT INTO tenants (name) VALUES 
('Sociomile Enterprise'),
('Kiki Tech Solutions');

-- seed users
-- password is "password123"
INSERT INTO users (id, tenant_id, name, email, password_hash, role) VALUES 
('admin-1', 1, 'Admin Sociomile', 'admin@sociomile.com', '$2a$10$ByI67Zgh.94.C8B/JtO.OuefN7dF3FjN/965B.35pYdJb6K4B5eC.', 'admin'),
('agent-1', 1, 'Agent Kiki', 'agent@sociomile.com', '$2a$10$ByI67Zgh.94.C8B/JtO.OuefN7dF3FjN/965B.35pYdJb6K4B5eC.', 'agent'),
('agent-2', 2, 'Agent Kiki', 'agent@tech.com', '$2a$10$ByI67Zgh.94.C8B/JtO.OuefN7dF3FjN/965B.35pYdJb6K4B5eC.', 'agent');
