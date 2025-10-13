# Setting Up RELEASE_TOKEN for Automatic Releases

The automatic release system requires a Personal Access Token (PAT) to trigger the release workflow when a tag is created. This is necessary because GitHub's default `GITHUB_TOKEN` cannot trigger other workflows (security feature to prevent infinite loops).

## Steps to Create and Add the PAT Token

### 1. Create a Personal Access Token

1. Go to GitHub Settings: https://github.com/settings/tokens
2. Click **"Generate new token"** → **"Generate new token (classic)"**
3. Give it a descriptive name: `FOF9 Editor Auto-Release Token`
4. Set expiration: Choose your preference (recommended: 1 year, then renew)
5. Select scopes:
   - ✅ **`repo`** (Full control of private repositories)
     - This includes: `repo:status`, `repo_deployment`, `public_repo`, `repo:invite`, `security_events`
6. Click **"Generate token"**
7. **IMPORTANT**: Copy the token immediately - you won't be able to see it again!

### 2. Add Token as Repository Secret

1. Go to your repository: https://github.com/igorilic/fof9editor
2. Click **Settings** tab
3. In the left sidebar, click **Secrets and variables** → **Actions**
4. Click **"New repository secret"**
5. Enter:
   - **Name**: `RELEASE_TOKEN`
   - **Value**: Paste the PAT token you copied
6. Click **"Add secret"**

### 3. Verify Setup

After adding the secret:

1. Make a change to CHANGELOG.md (e.g., add a new version section)
2. Commit and push to main
3. The auto-tag workflow should:
   - Create a new tag
   - Successfully trigger the release workflow
4. Check Actions tab to verify both workflows ran

## Token Permissions

The `RELEASE_TOKEN` needs the `repo` scope because it must:
- Push tags to the repository
- Trigger workflows (which requires write access to repository contents)

## Security Notes

- This token has write access to your repository
- Keep it secret - never commit it to code
- GitHub Actions secrets are encrypted and only exposed to workflows
- Consider setting an expiration date and renewing periodically
- You can revoke the token at any time from: https://github.com/settings/tokens

## Troubleshooting

### Token not working?
- Verify the token has `repo` scope
- Check that the secret name is exactly `RELEASE_TOKEN` (case-sensitive)
- Try regenerating the token if it's expired

### Workflow still not triggering?
- Check Actions tab for error messages
- Verify auto-tag workflow completed successfully
- Check that release.yml has correct trigger pattern: `tags: - 'v*.*.*'`

### Need to update the token?
1. Go to repository Settings → Secrets and variables → Actions
2. Click on `RELEASE_TOKEN`
3. Click **"Update secret"**
4. Paste new token value
5. Click **"Update secret"**

## Alternative: Fine-Grained Personal Access Token (Recommended for Security)

For better security, you can use a fine-grained PAT:

1. Go to: https://github.com/settings/tokens?type=beta
2. Click **"Generate new token"**
3. Token name: `FOF9 Editor Auto-Release`
4. Expiration: Your choice
5. Repository access: **"Only select repositories"** → `igorilic/fof9editor`
6. Permissions:
   - **Contents**: Read and write
   - **Metadata**: Read-only (automatically included)
   - **Workflows**: Read and write
7. Generate and copy token
8. Add as `RELEASE_TOKEN` secret (same process as above)

This approach limits the token to only this repository and only the permissions needed.
