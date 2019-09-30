import { writeFileSync, readFileSync, existsSync } from 'fs';

export class ConfigService {
  private file: string;
  private config: any;

  constructor(file: string) {
    this.file = file;
    if (!existsSync(file)) {
      writeFileSync(file, JSON.stringify({}));
    } else {
      this.config = JSON.parse(readFileSync(file).toString());
    }
  }

  get() {
    return this.config;
  }

  set(data: any) {
    writeFileSync(this.file, JSON.stringify(data));
    this.config = data;
  }
}
