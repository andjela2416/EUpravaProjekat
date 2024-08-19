export interface Appointment {
    studentId: string;
    date: Date;
    door_number: number;
    description: string;
    systematic: boolean;
    faculty_name:string;
    field_of_study:string;
    reserved:boolean;
  }

export interface TherapyData {
  id?: string;
  StudentHealthRecordId: string;
  Diagnosis: string;
}


