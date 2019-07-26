package protocol

const AuthCmdFormat = "*2\r\n$4\r\nauth\r\n$%d\r\n%s\r\n"
const SetCmdFormat = "*3\r\n$3\r\nset\r\n$%d\r\n%s\r\n$%d\r\n%s\r\n"
const GetCmdFormat = "*2\r\n$3\r\nget\r\n$%d\r\n%s\r\n"
const Newline = "\r\n"