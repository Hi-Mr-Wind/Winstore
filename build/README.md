# Build Directory

构建目录用于容纳应用程序的所有构建文件和资产。

结构是：

* bin - 输出目录
* darwin - macOS 特定文件
* windows - Windows 特定文件

## Mac

The `darwin` directory holds files specific to Mac builds.
These may be customised and used as part of the build. To return these files to the default state, simply delete them
and
build with `wails build`.

The directory contains the following files:

- `Info.plist` - the main plist file used for Mac builds. It is used when building using `wails build`.
- `Info.dev.plist` - same as the main plist file but used when building using `wails dev`.

## Windows

'windows' 目录包含使用 'wails build' 构建时使用的清单和 rc 文件。
这些可以根据您的应用进行定制。要将这些文件恢复到默认状态，只需删除它们并
使用 'wails build' 构建。

- `icon.ico` - 用于应用程序的图标。这在使用 'wails build' 进行构建时使用。如果您愿意
  使用不同的图标，只需将此文件替换为您自己的文件即可。如果缺少，则为新的“icon.ico”文件
  将使用构建目录中的`appicon.png`文件创建。
- `installer/*` -用于创建 Windows 安装程序的文件。这些是在使用 `wails build`.
- `info.json` - 用于 Windows 版本的应用程序详细信息。此处的数据将由 Windows 安装程序使用
  以及应用程序本身（右键单击 exe -> properties -> details）
- `wails.exe.manifest` - 主应用程序清单文件。