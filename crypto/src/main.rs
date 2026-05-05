use crypto_lib::{hash_password, secure_token, sha256_hex, verify_password};
use serde::{Deserialize, Serialize};
use std::io::{self, BufRead};

#[derive(Deserialize)]
#[serde(tag = "op", rename_all = "snake_case")]
enum Request {
    Hash { password: String },
    Verify { password: String, hash: String },
    Sha256 { input: String },
    Token,
}

#[derive(Serialize)]
#[serde(untagged)]
enum Response {
    Hash { hash: String },
    Verify { valid: bool },
    Sha256 { digest: String },
    Token { token: String },
    Error { error: String },
}

fn main() {
    let stdin = io::stdin();
    for line in stdin.lock().lines() {
        let line = match line {
            Ok(l) if !l.trim().is_empty() => l,
            _ => continue,
        };
        let resp = match serde_json::from_str::<Request>(&line) {
            Err(e) => Response::Error { error: e.to_string() },
            Ok(Request::Hash { password }) => match hash_password(&password) {
                Ok(hash) => Response::Hash { hash },
                Err(e) => Response::Error { error: e.to_string() },
            },
            Ok(Request::Verify { password, hash }) => match verify_password(&password, &hash) {
                Ok(valid) => Response::Verify { valid },
                Err(e) => Response::Error { error: e.to_string() },
            },
            Ok(Request::Sha256 { input }) => Response::Sha256 { digest: sha256_hex(&input) },
            Ok(Request::Token) => Response::Token { token: secure_token() },
        };
        println!("{}", serde_json::to_string(&resp).unwrap());
    }
}
