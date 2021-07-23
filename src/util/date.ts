export function dateBetween(start: Date, end: Date): Date {
  const millis = start.getTime() + (end.getTime() - start.getTime()) / 2;
  return new Date(millis);
}
