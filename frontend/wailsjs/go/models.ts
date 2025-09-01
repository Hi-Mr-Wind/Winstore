export namespace model {
	
	export class BaseModel {
	    code: number;
	    message: string;
	    data?: any;
	
	    static createFrom(source: any = {}) {
	        return new BaseModel(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}
	export class Collect {
	    id: string;
	    appName: string;
	    downloadUrl?: string;
	    appUrl?: string;
	    appIcon?: string;
	    createTime?: string;
	
	    static createFrom(source: any = {}) {
	        return new Collect(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.appName = source["appName"];
	        this.downloadUrl = source["downloadUrl"];
	        this.appUrl = source["appUrl"];
	        this.appIcon = source["appIcon"];
	        this.createTime = source["createTime"];
	    }
	}

}

export namespace options {
	
	export class SecondInstanceData {
	    Args: string[];
	    WorkingDirectory: string;
	
	    static createFrom(source: any = {}) {
	        return new SecondInstanceData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Args = source["Args"];
	        this.WorkingDirectory = source["WorkingDirectory"];
	    }
	}

}

export namespace sysUtils {
	
	export class SoftwareList {
	    Name: string;
	    Version: string;
	    Publisher: string;
	
	    static createFrom(source: any = {}) {
	        return new SoftwareList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Version = source["Version"];
	        this.Publisher = source["Publisher"];
	    }
	}
	export class WinSysInfo {
	    SysType: string;
	    SysVersion: string;
	    SysArch: string;
	    ComputerName: string;
	    TotalMemory: number;
	    CpuType: string;
	    CpuCoreNum: number;
	    CpuLogicNum: number;
	    Uptime: string;
	    LastBootUpTime: string;
	
	    static createFrom(source: any = {}) {
	        return new WinSysInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.SysType = source["SysType"];
	        this.SysVersion = source["SysVersion"];
	        this.SysArch = source["SysArch"];
	        this.ComputerName = source["ComputerName"];
	        this.TotalMemory = source["TotalMemory"];
	        this.CpuType = source["CpuType"];
	        this.CpuCoreNum = source["CpuCoreNum"];
	        this.CpuLogicNum = source["CpuLogicNum"];
	        this.Uptime = source["Uptime"];
	        this.LastBootUpTime = source["LastBootUpTime"];
	    }
	}

}

