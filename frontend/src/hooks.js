import { EventsOnce, EventsEmit } from "@wailsapp/runtime"

let unsubscribe;

export function openFileDialog() {
    unsubscribe = EventsOnce('openDirectoryDialog_result', (path) => {
        console.log('Directory selected:', path);
    });
}
