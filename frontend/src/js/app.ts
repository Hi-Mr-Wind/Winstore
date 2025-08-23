import {defineStore} from "pinia";
import {reactive} from "vue";
export const useAppStore = defineStore('app', ()=>{
    const map = reactive(new Map<string,any>())
    return {
        map,
    }
});