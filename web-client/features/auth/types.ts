export interface defaultResponse {
  type: string;
  message: string;
}

export interface AppJWTPayload {
  sub: string; // user's UUID
  email: string;
  first_name: string | null;
  last_name: string | null;
  picture: string | null;
  iat: number;
  exp: number;
  iss: string;
  aud: string;
}
