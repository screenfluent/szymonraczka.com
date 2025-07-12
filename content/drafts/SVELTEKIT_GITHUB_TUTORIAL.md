# SvelteKit + GitHub Tutorial: From Zero to Deployed

A complete guide to creating a new SvelteKit project and connecting it to GitHub.

## Prerequisites

### 1. Install Required Tools

```bash
# Install Bun (fast package manager)
curl -fsSL https://bun.sh/install | bash

# Install Git (if not already installed)
# macOS: git comes pre-installed
# Windows: https://git-scm.com/download/win
# Linux: sudo apt install git
```

### 2. GitHub Account Setup

1. **Create GitHub Account**: Visit [github.com](https://github.com) and sign up
2. **Verify Email**: Check your email and verify your account

### 3. SSH Key Setup for GitHub

```bash
# Generate SSH key (replace with your GitHub email)
ssh-keygen -t ed25519 -C "your-email@example.com"

# When prompted:
# - Press Enter for default file location
# - Enter a passphrase (optional but recommended)

# Start SSH agent
eval "$(ssh-agent -s)"

# Add SSH key to agent
ssh-add ~/.ssh/id_ed25519

# Copy public key to clipboard
# macOS:
pbcopy < ~/.ssh/id_ed25519.pub
# Linux:
cat ~/.ssh/id_ed25519.pub
# Windows (Git Bash):
clip < ~/.ssh/id_ed25519.pub
```

4. **Add SSH Key to GitHub**:
   - Go to GitHub → Settings → SSH and GPG keys
   - Click "New SSH key"
   - Paste your public key
   - Give it a descriptive title
   - Click "Add SSH key"

5. **Test SSH Connection**:
```bash
ssh -T git@github.com
# Should see: "Hi username! You've successfully authenticated..."
```

## Project Creation

### 1. Create SvelteKit Project

```bash
# Navigate to your projects directory
cd ~/Projects

# Create new SvelteKit project
bunx sv create your-project-name
```

**Configuration Options** (based on your example):
- Template: `SvelteKit minimal`
- TypeScript: `Yes, using TypeScript syntax`
- Add-ons: `prettier, eslint, tailwindcss, sveltekit-adapter, devtools-json, mdsvex`
- TailwindCSS plugins: `none`
- SvelteKit adapter: `netlify` (or your preferred deployment target)
- Package manager: `bun`

### 2. Initialize Git Repository

```bash
# Enter project directory
cd your-project-name

# Initialize git and make initial commit
git init && git add -A && git commit -m "Initial commit"

# Rename branch to main (modern standard)
git branch -M main
```

## GitHub Repository Setup

### 1. Create GitHub Repository

1. Go to [github.com/new](https://github.com/new)
2. Repository name: `your-project-name`
3. Description: Brief project description
4. Visibility: Public or Private
5. **Don't** initialize with README, .gitignore, or license (we already have these)
6. Click "Create repository"

### 2. Connect Local to Remote

```bash
# Add remote origin (replace with your repository URL)
git remote add origin git@github.com:yourusername/your-project-name.git

# Push to GitHub
git push -u origin main
```

## Development Workflow

### 1. Start Development Server

```bash
# Start dev server
bun run dev --open

# Server runs on http://localhost:5173
# --open flag automatically opens browser
```

### 2. Basic Git Workflow

```bash
# Check status
git status

# Stage changes
git add .

# Commit with descriptive message
git commit -m "feat: add homepage layout"

# Push to GitHub
git push
```

### 3. Useful Commands

```bash
# Install new dependencies
bun add package-name
bun add -d package-name  # dev dependency

# Build for production
bun run build

# Preview production build
bun run preview

# Run linting
bun run lint

# Format code
bun run format
```

## Project Structure

```
your-project-name/
├── src/
│   ├── lib/           # Reusable components
│   ├── routes/        # File-based routing
│   │   ├── +layout.svelte
│   │   └── +page.svelte
│   ├── app.html       # HTML template
│   ├── app.css        # Global styles
│   └── app.d.ts       # Type definitions
├── static/            # Static assets
├── package.json       # Dependencies
├── svelte.config.js   # Svelte configuration
├── vite.config.ts     # Vite configuration
└── tsconfig.json      # TypeScript config
```

## Deployment

### Netlify (if you chose netlify adapter)

1. Connect GitHub repository to Netlify
2. Build command: `bun run build`
3. Publish directory: `build`
4. Deploy automatically on push to main

### Vercel Alternative

```bash
# Switch to Vercel adapter
bun add -d @sveltejs/adapter-vercel

# Update svelte.config.js
# import adapter from '@sveltejs/adapter-vercel';
```

## Fixing Mistakes

### Changing Commit Messages

**Last commit not pushed yet** (most common case):
```bash
# Change the last commit message
git commit --amend -m "content: your new message"

# If you want to edit interactively
git commit --amend
```

**Last commit already pushed**:
```bash
# Change the message
git commit --amend -m "content: your new message"

# Force push (⚠️ use with caution on shared branches)
git push --force-with-lease
```

**Multiple commits back** (not pushed):
```bash
# Interactive rebase for last 3 commits
git rebase -i HEAD~3

# In the editor:
# - Change 'pick' to 'reword' for commits you want to change
# - Save and close
# - Edit each commit message when prompted
```

**Example: Changing "docs:" to "content:"**:
```bash
# Before: "docs: add tutorial for SvelteKit setup"
# After:  "content: add tutorial for SvelteKit setup"

git commit --amend -m "content: add tutorial for SvelteKit setup"
```

⚠️ **Important**: Only use `--force-with-lease` on branches you own. Never force push to shared branches like `main` without team coordination.

## Troubleshooting

### Common Issues

**Permission denied (publickey)**:
- Verify SSH key is added to GitHub
- Test with `ssh -T git@github.com`

**Branch 'main' doesn't exist**:
- Check current branch: `git branch`
- Rename if needed: `git branch -M main`

**HTTPS vs SSH**:
- Always use SSH for easier authentication
- Format: `git@github.com:username/repo.git`

## Best Practices

1. **Commit Messages**: Use conventional commits (`feat:`, `fix:`, `docs:`)
2. **Branch Protection**: Enable on GitHub for main branch
3. **Environment Variables**: Use `.env` files (add to `.gitignore`)
4. **Dependencies**: Keep `bun.lock` in version control
5. **Code Quality**: Set up pre-commit hooks with Husky

## Next Steps

- Set up GitHub Actions for CI/CD
- Add testing with Vitest
- Configure Playwright for E2E testing
- Set up issue templates
- Add contribution guidelines

---

**Pro Tip**: Keep it lean! Start minimal and add complexity only when needed. SvelteKit's philosophy aligns perfectly with building exactly what you need, nothing more.