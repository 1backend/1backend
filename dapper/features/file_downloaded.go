/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package features

import dt "github.com/openorch/openorch/dapper/types"

var FileDownloaded = dt.Feature{
	ID:   "file-downloaded",
	Name: "File Downloaded",
	Arguments: []dt.Argument{
		{
			Name:    "url",
			Type:    dt.String,
			Default: "",
		},
		{
			Name:    "assetfolder",
			Type:    dt.String,
			Default: "",
		},
		{
			Name:    "filename",
			Type:    dt.String,
			Default: "",
		},
	},
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Windows: {
			Execute: &dt.Script{
				Source: `# Define asset folder and filename
$assetfolder = "{{.assetfolder}}"
$filename = "{{.filename}}"
$fullPath = Join-Path -Path $assetfolder -ChildPath $filename
$backupPath = "$fullPath.bak"

# Function to download file with progress
function Download-FileWithProgress {
    param (
        [string]$url,
        [string]$destination
    )
    Write-Host "Starting file download, url: $url, destination: $destination"
    $webClient = New-Object System.Net.WebClient
    $webClient.DownloadFileAsync([Uri]$url, $destination)

    # Loop to track the download progress
    do {
        if (Test-Path $destination) {
            $size = (Get-Item $destination).length

            if ($webClient.ResponseHeaders -and $webClient.ResponseHeaders["Content-Length"]) {
                $response = $webClient.ResponseHeaders["Content-Length"]
                $totalSize = [int]$response
                $progress = [math]::Round(($size / $totalSize) * 100, 2)
                Write-Output "Downloading: $progress% Complete"
            } else {
                Write-Output "Downloading: Size unknown"
            }
        }
        Start-Sleep -Milliseconds 500 # Adjust the delay as needed
    } until ($webClient.IsBusy -eq $false)

    Write-Output "Download complete."
}

# Check if the file already exists
if (Test-Path -Path $fullPath) {
    # Get the length of the local file
    $localFileSize = (Get-Item $fullPath).length

    Write-Host "Get the length of the remote file"
    $response = Invoke-WebRequest -Uri "{{.url}}" -Method Head
    $remoteFileSize = $response.Headers["Content-Length"]
    
    if ($localFileSize -ne $remoteFileSize) {
        Write-Host "Backing up and deleting the local file because size doesn't match remote size"
        Copy-Item -Path $fullPath -Destination $backupPath
        Remove-Item -Path $fullPath
        Download-FileWithProgress -url "{{.url}}" -destination $fullPath
        Write-Host "File $filename was outdated and has been replaced in $assetfolder"
    } else {
        Write-Host "File $filename is up to date in $assetfolder"
    }
} else {
    # Download file if not present
    Download-FileWithProgress -url "{{.url}}" -destination $fullPath
    Write-Host "File $filename downloaded to $assetfolder"
}`,
				Runtime: dt.Powershell,
				Sudo:    false,
			},
			Check: &dt.Script{
				Source: `
# Define asset folder and filename
$assetfolder = "{{.assetfolder}}"
$filename = "{{.filename}}"
$fullPath = Join-Path -Path $assetfolder -ChildPath $filename

# Get the length of the remote file
$response = Invoke-WebRequest -Uri "{{.url}}" -Method Head
$remoteFileSize = $response.Headers["Content-Length"]

# Check if the file exists
if (Test-Path -Path $fullPath) {
    # Get the length of the local file
    $localFileSize = (Get-Item $fullPath).length
    
    if ($localFileSize -eq $remoteFileSize) {
        exit 0
    } else {
        exit 1
    }
} else {
    exit 1
}
`,
				Runtime: dt.Powershell,
				Sudo:    false,
			},
		},
		dt.Linux: {
			Execute: &dt.Script{
				Source: `
# Define asset folder and filename
assetfolder="{{.assetfolder}}"
filename="{{.filename}}"
fullPath="$assetfolder/$filename"
backupPath="$fullPath.bak"

# Get the length of the remote file
remoteFileSize=$(curl -sI "{{.url}}" | grep -i Content-Length | awk '{print $2}' | tr -d '\r')

# Check if the file already exists
if [ -f "$fullPath" ]; then
    # Get the length of the local file
    localFileSize=$(stat -c%s "$fullPath")
    
    if [ "$localFileSize" -ne "$remoteFileSize" ]; then
        # Backup and delete the local file if sizes don't match
        cp "$fullPath" "$backupPath"
        rm "$fullPath"
        wget -O "$fullPath" "{{.url}}"
        echo "File $filename was outdated and has been replaced in $assetfolder"
    else
        echo "File $filename is up to date in $assetfolder"
    fi
else
    # Download file if not present
    wget -O "$fullPath" "{{.url}}"
    echo "File $filename downloaded to $assetfolder"
fi
`,
				Runtime: dt.Bash,
				Sudo:    false,
			},
			Check: &dt.Script{
				Source: `
# Define asset folder and filename
assetfolder="{{.assetfolder}}"
filename="{{.filename}}"
fullPath="$assetfolder/$filename"

# Get the length of the remote file
remoteFileSize=$(curl -sI "{{.url}}" | grep -i Content-Length | awk '{print $2}' | tr -d '\r')

# Check if the file exists
if [ -f "$fullPath" ]; then
    # Get the length of the local file
    localFileSize=$(stat -c%s "$fullPath")
    
    if [ "$localFileSize" -eq "$remoteFileSize" ]; then
        exit 0
    else
        exit 1
    fi
else
    exit 1
fi
`,
				Runtime: dt.Bash,
				Sudo:    false,
			},
		},
		dt.MacOS: {
			Execute: &dt.Script{
				Source: `
#!/bin/bash

# Define asset folder and filename
assetfolder="{{.assetfolder}}"
filename="{{.filename}}"
fullPath="$assetfolder/$filename"
backupPath="$fullPath.bak"

# Get the length of the remote file
remoteFileSize=$(curl -sI "{{.url}}" | grep -i Content-Length | awk '{print $2}' | tr -d '\r')

# Check if the file already exists
if [ -f "$fullPath" ]; then
    # Get the length of the local file
    localFileSize=$(stat -f%z "$fullPath")
    
    if [ "$localFileSize" -ne "$remoteFileSize" ]; then
        # Backup and delete the local file if sizes don't match
        cp "$fullPath" "$backupPath"
        rm "$fullPath"
        curl -o "$fullPath" "{{.url}}"
        echo "File $filename was outdated and has been replaced in $assetfolder"
    else
        echo "File $filename is up to date in $assetfolder"
    fi
else
    # Download file if not present
    curl -o "$fullPath" "{{.url}}"
    echo "File $filename downloaded to $assetfolder"
fi
`,
				Runtime: dt.Bash,
				Sudo:    false,
			},
			Check: &dt.Script{
				Source: `
#!/bin/bash

# Define asset folder and filename
assetfolder="{{.assetfolder}}"
filename="{{.filename}}"
fullPath="$assetfolder/$filename"

# Get the length of the remote file
remoteFileSize=$(curl -sI "{{.url}}" | grep -i Content-Length | awk '{print $2}' | tr -d '\r')

# Check if the file exists
if [ -f "$fullPath" ]; then
    # Get the length of the local file
    localFileSize=$(stat -f%z "$fullPath")
    
    if [ "$localFileSize" -eq "$remoteFileSize" ]; then
        exit 0
    else
        exit 1
    fi
else
    exit 1
fi
`,
				Runtime: dt.Bash,
				Sudo:    false,
			},
		},
	},
	PlatformFeatures: map[dt.Platform][]any{},
}
