#!/bin/sh

# Attribution https://github.com/getseabird/seabird/blob/main/build/darwin/seabird.sh

cd "$(dirname "$0")" || exit
LAUNCH_DIR="$(pwd)"
export DYLD_LIBRARY_PATH="$LAUNCH_DIR/lib"
export GSETTINGS_SCHEMA_DIR="$LAUNCH_DIR/share/glib-2.0/schemas"
export GDK_PIXBUF_MODULEDIR="$LAUNCH_DIR/lib/gdk-pixbuf-2.0"
export GDK_PIXBUF_MODULE_FILE="$LAUNCH_DIR/lib/gdk-pixbuf-2.0/2.10.0/loaders.cache"
export XDG_DATA_DIRS="$LAUNCH_DIR/share"

export PATH="$PATH:/opt/homebrew/bin:/opt/local/bin"

"$LAUNCH_DIR/rokon-gtk" "$@"
