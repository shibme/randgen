#!/usr/bin/env pwsh

$ErrorActionPreference = 'Stop'

if ($v) {
  $Version = "v${v}"
}
if ($Args.Length -eq 1) {
  $Version = $Args.Get(0)
}

$RandGenInstall = $env:RANDGEN_INSTALL
$BinDir = if ($RandGenInstall) {
  "${RandGenInstall}\bin"
} else {
  "${Home}\.randgen\bin"
}

$RandGenZip = "$BinDir\randgen.zip"
$RandGenExe = "$BinDir\randgen.exe"
$Target = 'windows_amd64'

$DownloadUrl = if (!$Version) {
  "https://github.com/shibme/randgen/releases/latest/download/randgen_${Target}.zip"
} else {
  "https://github.com/shibme/randgen/releases/download/${Version}/randgen_${Target}.zip"
}

if (!(Test-Path $BinDir)) {
  New-Item $BinDir -ItemType Directory | Out-Null
}

curl.exe -Lo $RandGenZip $DownloadUrl

tar.exe xf $RandGenZip -C $BinDir

Remove-Item $RandGenZip

$User = [System.EnvironmentVariableTarget]::User
$Path = [System.Environment]::GetEnvironmentVariable('Path', $User)
if (!(";${Path};".ToLower() -like "*;${BinDir};*".ToLower())) {
  [System.Environment]::SetEnvironmentVariable('Path', "${Path};${BinDir}", $User)
  $Env:Path += ";${BinDir}"
}

Write-Output "RandGen was installed successfully to ${RandGenExe}"
Write-Output "Run 'randgen --help' to get started"
