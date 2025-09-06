# Mini ERP
Go + PostgreSQL + Vue + Docker 연습용 프로젝트
<br> 인증, 프로젝트 관리, 작업 관리, 간단한 파일 업로드/버전 관리

---

## [기능]
- 회원가입 / 로그인 (JWT 인증)
- 프로젝트 관리 (생성, 조회, 삭제)
- 작업 관리 (TODO, INPROGRESS, COMPLETED)
- 파일 업로드 및 버전 관리 (동일 파일명 업로드 시 자동 버전 증가)

---

## [사용 기술]
- **Backend**: Go (Gin, GORM, JWT, Bcrypt)
- **DB**: PostgreSQL
- **Frontend**: Vue.js (기본 UI)
- **Infra**: Docker, Docker Compose

---

## 실행 방법

### 1. 저장소 클론
```bash
git clone https://github.com/<user>/mini-erp.git
cd mini-erp
