export interface Location {
    id: number;
    name: string;
    type?: string;
    dimension?: string;
    url?: string;
    created?: string;
  }

  export function validateLocation(location: Location | undefined): boolean {
    if (!location) return true; 
  
    if (typeof location.id !== 'number' || location.id <= 0) {
      return false;
    }
  
    if (typeof location.name !== 'string' || location.name.trim() === '') {
      return false;
    }
  
  
    return true;  
  }